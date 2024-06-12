// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useEffect, useState, useRef} from 'react';
import {FormattedMessage} from 'react-intl';
import {useSelector, useDispatch} from 'react-redux';
import styled from 'styled-components';

import {getCurrentChannelNameForSearchShortcut} from 'mattermost-redux/selectors/entities/channels';

import {
    updateSearchTerms,
    showSearchResults,
    updateSearchType,
} from 'actions/views/rhs';
import {
    getSearchTerms,
} from 'selectors/rhs';

import Popover from 'components/widgets/popover';

import Constants from 'utils/constants';
import * as Keyboard from 'utils/keyboard';
import {isServerVersionGreaterThanOrEqualTo} from 'utils/server_version';
import {isDesktopApp, getDesktopVersion, isMacApp} from 'utils/user_agent';

import type {GlobalState} from 'types/store';

import SearchBox from './search_box';

type Props = {
    enableFindShortcut: boolean;
}

const PopoverStyled = styled(Popover)`
    min-width: 600px;
    left: -90px;
    top: -12px;
    border-radius: 12px;

    .popover-content {
        padding: 0px;
    }
`;

const NewSearchContainer = styled.div`
    display: flex;
    position: relative;
    align-items: center;
    height: 28px;
    width: 100%;
    background-color: rgba(var(--sidebar-text-rgb), 0.08);
    color: rgba(var(--sidebar-text-rgb), 0.64);
    font-size: 12px;
    font-weight: 500;
    border-radius: var(--radius-s);
    border: none;
    padding: 4px;
    cursor: pointer;
    &:hover {
        background-color: rgba(var(--sidebar-text-rgb), 0.16);
        color: rgba(var(--sidebar-text-rgb), 0.88);
    }
`;

const NewSearch = ({enableFindShortcut}: Props): JSX.Element => {
    const currentChannelName = useSelector(getCurrentChannelNameForSearchShortcut);
    const searchTerms = useSelector(getSearchTerms) || '';
    const pluginSearch = useSelector((state: GlobalState) => state.plugins.components.SearchButtons);
    const dispatch = useDispatch();
    const [focused, setFocused] = useState<boolean>(false);
    const [currentChannel, setCurrentChannel] = useState('');
    const searchBoxRef = useRef<HTMLDivElement|null>(null);

    useEffect(() => {
        if (!enableFindShortcut) {
            return undefined;
        }

        const isDesktop = isDesktopApp() && isServerVersionGreaterThanOrEqualTo(getDesktopVersion(), '4.7.0');

        const handleKeyDown = (e: KeyboardEvent) => {
            if (Keyboard.isKeyPressed(e, Constants.KeyCodes.ESCAPE)) {
                e.preventDefault();
                setCurrentChannel('');
                setFocused(false);
            }

            if (Keyboard.cmdOrCtrlPressed(e) && Keyboard.isKeyPressed(e, Constants.KeyCodes.F)) {
                if (!isDesktop && !e.shiftKey) {
                    return;
                }

                // Special case for Mac Desktop xApp where Ctrl+Cmd+F triggers full screen view
                if (isMacApp() && e.ctrlKey) {
                    return;
                }

                e.preventDefault();
                setCurrentChannel(currentChannelName || '');
                setFocused(true);
            }
        };

        document.addEventListener('keydown', handleKeyDown);
        return () => {
            document.removeEventListener('keydown', handleKeyDown);
        };
    }, [currentChannelName]);

    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (searchBoxRef.current) {
                if (e.target !== searchBoxRef.current && !searchBoxRef.current.contains(e.target as Node)) {
                    setFocused(false);
                    setCurrentChannel('');
                }
            }
        };

        document.addEventListener('click', handleClick, {capture: true});
        return () => {
            document.removeEventListener('click', handleClick);
        };
    }, []);

    return (
        <NewSearchContainer
            tabIndex={0}
            onKeyDown={(e: React.KeyboardEvent) => {
                if (Keyboard.isKeyPressed(e, Constants.KeyCodes.TAB)) {
                    return;
                }
                setFocused(true);
            }}
            onClick={() => setFocused(true)}
        >
            <i className='icon icon-magnify'/>
            {searchTerms && <span tabIndex={0}>{searchTerms}</span>}
            {!searchTerms && (
                <FormattedMessage
                    id='search_bar.search'
                    defaultMessage='Search'
                />
            )}
            {focused && (
                <PopoverStyled
                    id='searchPopover'
                    placement='bottom'
                >
                    <SearchBox
                        ref={searchBoxRef}
                        onClose={() => {
                            setFocused(false);
                            setCurrentChannel('');
                        }}
                        onSearch={(searchType: string, searchTerms: string) => {
                            dispatch(updateSearchType(searchType));
                            dispatch(updateSearchTerms(searchTerms));

                            if (searchType === '' || searchType === 'messages' || searchType === 'files' || searchType === 'omnisearch') {
                                dispatch(showSearchResults(false));
                            } else {
                                pluginSearch.forEach((pluginData: any) => {
                                    if (pluginData.pluginId === searchType) {
                                        pluginData.action(searchTerms);
                                    }
                                });
                            }
                            setFocused(false);
                            setCurrentChannel('');
                        }}
                        initialSearchTerms={currentChannel ? `in:${currentChannel} ` : searchTerms}
                    />
                </PopoverStyled>
            )}
        </NewSearchContainer>
    );
};

export default NewSearch;
