// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React, {memo, useCallback} from 'react';
import type {ReactNode} from 'react';
import {FormattedMessage} from 'react-intl';
import {useSelector, useDispatch} from 'react-redux';

import type {UserProfile} from '@mattermost/types/users';

import {getCurrentChannel, getCurrentChannelStats} from 'mattermost-redux/selectors/entities/channels';
import {getConfig} from 'mattermost-redux/selectors/entities/general';

import {
    showPinnedPosts,
    showChannelFiles,
    closeRightHandSide,
} from 'actions/views/rhs';
import {getRhsState} from 'selectors/rhs';

import {Constants, RHSStates} from 'utils/constants';
import {isFileAttachmentsEnabled as isFileAttachmentsEnabledCheck} from 'utils/file_utils';

import ChannelHeaderDmStatus from './channel_header_dm_status';
import ChannelHeaderEditMessage from './channel_header_edit_message';
import ChannelHeaderMembersButton from './channel_header_members_button';
import ChannelHeaderPopover from './channel_header_popover';
import HeaderIconWrapper from './components/header_icon_wrapper';

type Props = {
    dmUser?: UserProfile;
}

const ChannelHeaderText = ({
    dmUser,
}: Props) => {
    const config = useSelector(getConfig);
    const isFileAttachmentsEnabled = isFileAttachmentsEnabledCheck(config);
    const hideGuestTags = config.HideGuestTags === 'true';
    const dispatch = useDispatch();
    const channel = useSelector(getCurrentChannel) || {};
    const isDirect = (channel.type === Constants.DM_CHANNEL);
    const isGroup = (channel.type === Constants.GM_CHANNEL);
    const rhsState = useSelector(getRhsState);
    const stats = useSelector(getCurrentChannelStats) || EMPTY_CHANNEL_STATS;
    const hasGuests = stats.guest_count > 0;
    const pinnedPostsCount = stats.pinnedpost_count;

    const showPinnedPostsCallback = useCallback((e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        if (rhsState === RHSStates.PIN) {
            dispatch(closeRightHandSide());
        } else {
            dispatch(showPinnedPosts());
        }
    }, [rhsState]);

    const showChannelFilesCallback = useCallback(() => {
        if (rhsState === RHSStates.CHANNEL_FILES) {
            dispatch(closeRightHandSide());
        } else {
            dispatch(showChannelFiles(channel.id));
        }
    }, [rhsState, channel.id]);

    const channelFilesIconClass = classNames('channel-header__icon channel-header__icon--wide channel-header__icon--left', {
        'channel-header__icon--active': rhsState === RHSStates.CHANNEL_FILES,
    });
    const channelFilesIcon = <i className='icon icon-file-text-outline'/>;
    const pinnedIconClass = classNames('channel-header__icon channel-header__icon--wide channel-header__icon--left', {
        'channel-header__icon--active': rhsState === RHSStates.PIN,
    });

    const pinnedIcon = (
        <>
            <i
                aria-hidden='true'
                className='icon icon-pin-outline channel-header__pin'
            />
            {pinnedPostsCount &&
                <span
                    id='channelPinnedPostCountText'
                    className='icon__text'
                >
                    {pinnedPostsCount}
                </span>
            }
        </>
    );

    let hasGuestsText: ReactNode = '';
    if (hasGuests && !hideGuestTags) {
        hasGuestsText = (
            <span className='has-guest-header'>
                <span tabIndex={0}>
                    <FormattedMessage
                        id='channel_header.channelHasGuests'
                        defaultMessage='This channel has guests'
                    />
                </span>
            </span>
        );
    }

    if (isGroup) {
        if (hasGuests && !hideGuestTags) {
            hasGuestsText = (
                <span className='has-guest-header'>
                    <FormattedMessage
                        id='channel_header.groupMessageHasGuests'
                        defaultMessage='This group message has guests'
                    />
                </span>
            );
        }
    }

    let headerTextContainer;
    const headerText = (isDirect && dmUser?.is_bot) ? dmUser.bot_description : channel.header;

    if (headerText) {
        headerTextContainer = (
            <div
                id='channelHeaderDescription'
                className='channel-header__description'
                dir='auto'
            >
                <ChannelHeaderDmStatus dmUser={dmUser}/>
                <ChannelHeaderMembersButton/>

                <HeaderIconWrapper
                    iconComponent={pinnedIcon}
                    ariaLabel={true}
                    buttonClass={pinnedIconClass}
                    buttonId={'channelHeaderPinButton'}
                    onClick={showPinnedPostsCallback}
                    tooltipKey={'pinnedPosts'}
                />
                {isFileAttachmentsEnabled &&
                    <HeaderIconWrapper
                        iconComponent={channelFilesIcon}
                        ariaLabel={true}
                        buttonClass={channelFilesIconClass}
                        buttonId={'channelHeaderFilesButton'}
                        onClick={showChannelFilesCallback}
                        tooltipKey={'channelFiles'}
                    />
                }
                {hasGuestsText}
                <ChannelHeaderPopover dmUser={dmUser}/>
            </div>
        );
    } else {
        headerTextContainer = (
            <div
                id='channelHeaderDescription'
                className='channel-header__description light'
            >
                <ChannelHeaderDmStatus dmUser={dmUser}/>
                <ChannelHeaderMembersButton/>

                <HeaderIconWrapper
                    iconComponent={pinnedIcon}
                    ariaLabel={true}
                    buttonClass={pinnedIconClass}
                    buttonId={'channelHeaderPinButton'}
                    onClick={showPinnedPostsCallback}
                    tooltipKey={'pinnedPosts'}
                />
                {isFileAttachmentsEnabled &&
                    <HeaderIconWrapper
                        iconComponent={channelFilesIcon}
                        ariaLabel={true}
                        buttonClass={channelFilesIconClass}
                        buttonId={'channelHeaderFilesButton'}
                        onClick={showChannelFilesCallback}
                        tooltipKey={'channelFiles'}
                    />
                }
                {hasGuestsText}
                <ChannelHeaderEditMessage dmUser={dmUser}/>
            </div>
        );
    }
    return headerTextContainer;
};

export default memo(ChannelHeaderText);
