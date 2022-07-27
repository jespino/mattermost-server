//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	pipe "gopkg.in/pipe.v2"
)

type Docker mg.Namespace
type HAServer mg.Namespace
type Plugins mg.Namespace

func goBin() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "bin"), nil
}

func goBinRun(cmd string, args ...string) error {
	goBinPath, err := goBin()
	if err != nil {
		return err
	}

	return sh.Run(filepath.Join(goBinPath, cmd), args...)
}

func goRun(args ...string) error {
	if os.Getenv("GOFLAGS") != "" {
		arguments := append([]string{"run", os.Getenv("GOFLAGS")}, args...)
		return sh.Run("go", arguments...)
	}
	arguments := append([]string{"run"}, args...)
	return sh.Run("go", arguments...)
}

func goInstall(args ...string) error {
	arguments := append([]string{"install"}, args...)
	return sh.Run("go", arguments...)
}

func isM1() bool {
	fmt.Println(runtime.GOARCH, runtime.GOOS)
	return runtime.GOARCH == "arm64" && runtime.GOOS == "darwin"
}

func isEnterpriseReady() bool {
	return fileExists("../enterprise")
}

func isBoardsReady() bool {
	return fileExists("../focalboard")
}

func fileExists(path string) bool {
	if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func getDockerComposeOverride() []string {
	result := []string{}
	if isM1() {
		result = append(result, "-f", "docker-compose.makefile.m1.yml")
	}

	if fileExists("docker-compose.override.yaml") {
		result = append(result, "-f", "docker-compose.override.yaml")
	}

	return result
}

func (HAServer) Run() error {
	if !isEnterpriseReady() {
		return errors.New("Unable to run high availability server without enterprise code")
	}
	fmt.Println("Starting mattermost in an HA topology '(3 node cluster)'")
	arguments := append([]string{"-f", "docker-compose.yaml"}, getDockerComposeOverride()...)
	arguments = append(arguments, "up", "--remove-orphans", "haproxy")
	return sh.Run("docker-compose", arguments...)
}

func (HAServer) Stop() error {
	fmt.Println("Stopping docker containers for HA topology")
	return sh.Run("docker-compose", "stop")
}

func (Docker) Stop() error {
	if os.Getenv("MM_NO_DOCKER") == "true" {
		fmt.Println("No docker enabled: skipping docker stop")
		return nil
	}
	fmt.Println("Stopping docker containers")
	return sh.Run("docker-compose", "stop")
}

func (Docker) Clean() error {
	if os.Getenv("MM_NO_DOCKER") == "true" {
		fmt.Println("No docker enabled: skipping docker clean")
		return nil
	}
	fmt.Println("Removing docker containers")
	err := sh.Run("docker-compose", "down", "-v")
	if err != nil {
		return err
	}
	return sh.Run("docker-compose", "rm", "-v")

}

func (Docker) Start() error {
	if os.Getenv("IS_CI") == "true" {
		fmt.Println("CI Build: skipping docker start")
		return nil
	}

	if os.Getenv("MM_NO_DOCKER") == "true" {
		fmt.Println("No Docker Enabled: skipping docker start")
		return nil
	}

	fmt.Println("Starting docker containers")

	err := sh.Run("docker-compose", "rm", "start_dependencies")
	if err != nil {
		return err
	}

	services := strings.Fields(os.Getenv("ENABLED_DOCKER_SERVICES"))
	arguments := append([]string{"run", "./build/docker-compose-generator/main.go"}, services...)
	dockerCompose, err := sh.Output("go", arguments...)
	if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.WriteString(dockerCompose)

	arguments = append([]string{"-f", "docker-compose.makefile.yml", "-f", tmpfile.Name()}, getDockerComposeOverride()...)
	arguments = append(arguments, "run", "-T", "--rm", "start_dependencies")
	err = sh.Run("docker-compose", arguments...)
	if err != nil {
		return err
	}

	if strings.Contains("openldap", os.Getenv("ENABLED_DOCKER_SERVICES")) {
		arguments = append([]string{"-f", "docker-compose.makefile.yml"}, getDockerComposeOverride()...)
		arguments = append(arguments, "exec", "-T", "openldap", "bash", "-c", "ldapadd -x -D \"cn=admin,dc=mm,dc=test,dc=com\" -w mostest || true")
		ldapDataFile := "tests/test-data.ldif"
		if os.Getenv("LDAP_DATA") != "" {
			ldapDataFile = "tests/" + os.Getenv("LDAP_DATA") + "-data.ldif"
		}
		err = pipe.Run(pipe.Line(
			pipe.ReadFile(ldapDataFile),
			pipe.Exec("docker-compose", arguments...),
		))
		if err != nil {
			return err
		}
	}
	if strings.Contains("mysql-read-replica", os.Getenv("ENABLED_DOCKER_SERVICES")) {
		err = sh.Run("sh", "./scripts/replica-mysql-config.sh")
		if err != nil {
			return err
		}
	}
	return nil
}

var Default = Build

// Clean plus removes persistent server data.
func Nuke() error {
	fmt.Println("BOOM")
	mg.Deps(Clean, Docker.Clean)
	err := sh.Run("rm", "-rf", "data")
	if err != nil {
		return err
	}
	return sh.Run("rm", "-f", "go.work", "go.work.sum")
}

// Clean cleans up everything except persistent server data.
func Clean() error {
	fmt.Println("Cleaning...")

	sh.Run("rm", "-Rf", "$(DIST_ROOT)")
	sh.Run("go", "clean", os.Getenv("GOFLAGS"), "-i", "./...")

	os.Chdir("$(BUILD_WEBAPP_DIR)")
	sh.Run("$(MAKE)", "clean")
	os.Chdir("..")

	sh.Run("find", ".", "-type", "d", "-name", "data", "|", "xargs", "rm", "-rf")
	sh.Run("rm", "-rf", "logs")

	sh.Run("rm", "-f", "mattermost.log")
	sh.Run("rm", "-f", "mattermost.log.jsonl")
	sh.Run("rm", "-f", "npm-debug.log")
	sh.Run("rm", "-f", ".prepare-go")
	sh.Run("rm", "-f", "enterprise")
	sh.Run("rm", "-f", "cover.out")
	sh.Run("rm", "-f", "ecover.out")
	sh.Run("rm", "-f", "*.out")
	sh.Run("rm", "-f", "*.test")
	sh.Run("rm", "-f", "imports/imports.go")
	sh.Run("rm", "-f", "cmd/mattermost/cprofile*.out")
	return nil
}

func CheckStyle() error {
	mg.Deps(GolangCILint, Plugins.Check, GoVet)
	return nil
}

func GolangCILint() error {
	err := goInstall("github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2")
	if err != nil {
		return err
	}

	fmt.Println("Running golangci-lint")
	err = goBinRun("golangci-lint", "run", "./...")
	if err != nil {
		return err
	}
	if isEnterpriseReady() && os.Getenv("MM_NO_ENTERPRISE_LINT") != "true" {
		err = goBinRun("golangci-lint", "run", "../enterprise/...")
		if err != nil {
			return err
		}
	}
	if isBoardsReady() && os.Getenv("MM_NO_BOARDS_LINT") != "true" {
		err = goBinRun("golangci-lint", "run", "../focalboard/server/...")
		if err != nil {
			return err
		}
	}
	return nil
}

func GoVet() error {
	goBinPath, err := goBin()
	if err != nil {
		return err
	}
	err = goInstall("github.com/mattermost/mattermost-govet/v2@new")
	if err != nil {
		return err
	}
	arguments := []string{
		"vet", "-vettool=" + goBinPath + "/mattermost-govet",
		"-license", "-structuredLogging", "-inconsistentReceiverName",
		"-inconsistentReceiverName.ignore=session_serial_gen.go,team_member_serial_gen.go,user_serial_gen.go",
		"-emptyStrCmp", "-tFatal", "-configtelemetry", "-errorAssertions",
	}
	if os.Getenv("MM_VET_OPENSPEC_PATH") != "" && fileExists(os.Getenv("MM_VET_OPENSPEC_PATH")) {
		arguments = append(arguments, "-openApiSync", "-openApiSync.spec="+os.Getenv("MM_VET_OPENSPEC_PATH"))
	} else {
		fmt.Println("MM_VET_OPENSPEC_PATH not set or spec yaml path in it is incorrect. Skipping API check")
	}
	arguments = append(arguments, "./...")
	sh.Run("go", arguments...)
	if isEnterpriseReady() && os.Getenv("MM_NO_ENTERPRISE_LINT") != "true" {
		return sh.Run("go", "vet", "-vettool="+goBinPath+"/mattermost-govet", "-enterpriseLicense", "-structuredLogging", "-tFatal", "../enterprise/...")
	}
	return nil
}

func (Plugins) Check() error {
	return goRun(os.Getenv("GOFLAGS"), "./plugin/checker")
}

func (Plugins) Prepackaged() error {
	plugins := []string{
		"mattermost-plugin-antivirus-v0.1.2",
		"mattermost-plugin-autolink-v1.2.2",
		"mattermost-plugin-aws-SNS-v1.2.0",
		"mattermost-plugin-calls-v0.7.0",
		"mattermost-plugin-channel-export-v1.0.0",
		"mattermost-plugin-custom-attributes-v1.3.0",
		"mattermost-plugin-github-v2.0.1",
		"mattermost-plugin-gitlab-v1.3.0",
		"mattermost-plugin-playbooks-v1.29.1",
		"mattermost-plugin-jenkins-v1.1.0",
		"mattermost-plugin-jira-v2.4.0",
		"mattermost-plugin-nps-v1.2.0",
		"mattermost-plugin-welcomebot-v1.2.0",
		"mattermost-plugin-zoom-v1.6.0",
		"focalboard-v7.1.0",
		"mattermost-plugin-apps-v1.1.0",
	}
	fmt.Println("Downloading prepackaged plugins")
	err := sh.Run("mkdir", "-p", "prepackaged_plugins")
	if err != nil {
		return err
	}
	os.Chdir("prepackaged_plugins")
	defer os.Chdir("..")

	for _, plugin := range plugins {
		err = sh.Run("curl", "-f", "-O", "-L", fmt.Sprintf("https://plugins-store.test.mattermost.com/release/%s.tar.gz", plugin))
		if err != nil {
			return err
		}
		err = sh.Run("curl", "-f", "-O", "-L", fmt.Sprintf("https://plugins-store.test.mattermost.com/release/%s.tar.gz.sig", plugin))
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateGoVersion() error {
	minimumGoMajorVersion := "1"
	minimumGoMinorVersion := "15"
	versionString, err := sh.Output("go", "version")
	if err != nil {
		return err
	}
	version := strings.Fields(versionString)[2]
	versionSlice := strings.Split(version[2:], ".")
	majorVersion := versionSlice[0]
	minorVersion := versionSlice[1]
	if minimumGoMajorVersion < majorVersion {
		return nil
	}
	if minimumGoMajorVersion == majorVersion {
		if minimumGoMinorVersion <= minorVersion {
			return nil
		}
	}
	return fmt.Errorf("Golang version is not supported, please update to at least %s.%s", minimumGoMajorVersion, minimumGoMinorVersion)
}

func Build() error {
	mg.Deps(BuildLinuxAMD64, BuildLinuxARM64, BuildMacOsXAMD64, BuildMacOsXARM64, BuildWindows)
	return nil
}

func BuildLinux() error {
	mg.Deps(BuildLinuxAMD64, BuildLinuxARM64)
	return nil
}

func BuildLinuxAMD64() error {
	fmt.Println("Build Linux amd64")

	goBinPath, err := goBin()
	if err != nil {
		return err
	}

	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "amd64",
	}
	return sh.RunWith(env, "go", "build", "-o", goBinPath, "$(GOFLAGS)", "-trimpath", "-ldflags", "'$(LDFLAGS)'", "./...")
}

func BuildLinuxARM64() error {
	fmt.Println("Build Linux arm64")

	goBinPath, err := goBin()
	if err != nil {
		return err
	}

	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "arm64",
	}
	return sh.RunWith(env, "go", "build", "-o", goBinPath, "$(GOFLAGS)", "-trimpath", "-ldflags", "'$(LDFLAGS)'", "./...")
}

func BuildMacOsX() error {
	mg.Deps(BuildMacOsXAMD64, BuildMacOsXARM64)
	return nil
}

func BuildMacOsXAMD64() error {
	// TODO:
	return nil
}

func BuildMacOsXARM64() error {
	// TODO:
	return nil
}

func BuildWindows() error {
	// TODO:
	return nil
}
