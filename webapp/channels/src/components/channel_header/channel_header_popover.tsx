// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo, useState, useMemo, useCallback, useRef} from 'react';
import {useSelector} from 'react-redux';
import classNames from 'classnames';
import {Overlay} from 'react-bootstrap';

import type {UserProfile} from '@mattermost/types/users';

import {handleFormattedTextClick} from 'utils/utils';

import Markdown from 'components/markdown';
import Popover from 'components/widgets/popover';
import {getCurrentChannel} from 'mattermost-redux/selectors/entities/channels';
import {Constants} from 'utils/constants';
import {getCurrentRelativeTeamUrl, getMyTeams} from 'mattermost-redux/selectors/entities/teams';
import {getAnnouncementBarCount} from 'selectors/views/announcement_bar';

const headerMarkdownOptionsBase = {singleline: true, mentionHighlight: false, atMentions: true};
const popoverMarkdownOptionsBase = {singleline: false, mentionHighlight: false, atMentions: true};

type Props = {
    dmUser?: UserProfile;
}

const ChannelHeaderText = ({
    dmUser,
}: Props) => {
    const [topOffset, setTopOffset] = useState(0);
    const [leftOffset, setLeftOffset] = useState(0);
    const [channelHeaderPoverWidth, setChannelHeaderPoverWidth] = useState(0);
    const [showChannelHeaderPopover, setShowChannelHeaderPopover] = useState(false);
    const teams = useSelector(getMyTeams);
    const hasMoreThanOneTeam = teams.length > 1;
    const channel = useSelector(getCurrentChannel) || {};
    const isDirect = (channel.type === Constants.DM_CHANNEL);
    const headerDescriptionRef = useRef<HTMLSpanElement>(null);
    const headerOverlayRef = useRef<Overlay>(null);
    const channelNamesMap = channel.props && channel.props.channel_mentions;
    const headerMarkdownOptions = useMemo(() => ({...headerMarkdownOptionsBase, channelNamesMap}), [channelNamesMap]);
    const popoverMarkdownOptions = useMemo(() => ({...popoverMarkdownOptionsBase, channelNamesMap}), [channelNamesMap]);
    const headerPopoverTextMeasurerRef = useRef<HTMLDivElement>(null);
    const announcementBarCount = useSelector(getAnnouncementBarCount);
    const currentRelativeTeamUrl = useSelector(getCurrentRelativeTeamUrl);

    const handleFormattedTextClickCallback = useCallback((e: React.MouseEvent<HTMLSpanElement>) => handleFormattedTextClick(e, currentRelativeTeamUrl), [currentRelativeTeamUrl]);

    const showChannelHeaderPopoverCallback = useCallback((headerText: string) => {
        const headerDescriptionRect = headerDescriptionRef.current?.getBoundingClientRect();
        const headerPopoverTextMeasurerRect = headerPopoverTextMeasurerRef.current?.getBoundingClientRect();
        const announcementBarSize = 40;

        if (headerPopoverTextMeasurerRect && headerDescriptionRect) {
            if (headerPopoverTextMeasurerRect.width > headerDescriptionRect.width || headerText.match(/\n{2,}/g)) {
                const leftOffset = headerDescriptionRect.left - (hasMoreThanOneTeam ? 313 : 248);
                setShowChannelHeaderPopover(true);
                setLeftOffset(leftOffset);
            }
        }

        // add 40px to take the global header into account
        const topOffset = (announcementBarSize * announcementBarCount) + 40;
        const channelHeaderPoverWidth = headerDescriptionRef.current?.clientWidth || 0 - (hasMoreThanOneTeam ? 64 : 0);

        setTopOffset(topOffset);
        setChannelHeaderPoverWidth(channelHeaderPoverWidth);
    }, [hasMoreThanOneTeam, announcementBarCount]);

    const headerText = (isDirect && dmUser?.is_bot) ? dmUser.bot_description : channel.header;
    if (!headerText) {
        return null;
    }

    const imageProps = {
        hideUtilities: true,
    };
    const popoverContent = (
        <Popover
            id='header-popover'
            popoverStyle='info'
            popoverSize='lg'
            style={{transform: `translate(${leftOffset}px, ${topOffset}px)`, maxWidth: channelHeaderPoverWidth}}
            placement='bottom'
            className={classNames('channel-header__popover', {'chanel-header__popover--lhs_offset': hasMoreThanOneTeam})}
        >
            <span
                onClick={handleFormattedTextClickCallback}
            >
                <Markdown
                    message={headerText}
                    options={popoverMarkdownOptions}
                    imageProps={imageProps}
                />
            </span>
        </Popover>
    );

    return (
        <>
            <div
                className='header-popover-text-measurer'
                ref={headerPopoverTextMeasurerRef}
            >
                <Markdown
                    message={headerText.replace(/\n+/g, ' ')}
                    options={headerMarkdownOptions}
                    imageProps={imageProps}
                />
            </div>
            <span
                className='header-description__text'
                onClick={handleFormattedTextClickCallback}
                onMouseOver={() => showChannelHeaderPopoverCallback(headerText)}
                onMouseOut={() => setShowChannelHeaderPopover(false)}
                ref={headerDescriptionRef}
            >
                <Overlay
                    show={showChannelHeaderPopover}
                    placement='bottom'
                    rootClose={true}
                    target={headerDescriptionRef.current as React.ReactInstance}
                    ref={headerOverlayRef}
                    onHide={() => setShowChannelHeaderPopover(false)}
                >
                    {popoverContent}
                </Overlay>

                <Markdown
                    message={headerText}
                    options={headerMarkdownOptions}
                    imageProps={imageProps}
                />
            </span>
        </>
    )
}

export default memo(ChannelHeaderText);
