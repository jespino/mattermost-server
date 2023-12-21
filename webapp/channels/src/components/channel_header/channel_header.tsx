// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo, useEffect, useCallback} from 'react';
import {useSelector, useDispatch} from 'react-redux';
import {FormattedMessage, useIntl} from 'react-intl';

import {getCurrentChannel, isCurrentChannelMuted, getMyCurrentChannelMembership} from 'mattermost-redux/selectors/entities/channels';
import {getCurrentUser, getUser, makeGetProfilesInChannel} from 'mattermost-redux/selectors/entities/users';
import {getCustomEmojisInText} from 'mattermost-redux/actions/emojis';
import {updateChannelNotifyProps} from 'mattermost-redux/actions/channels';
import {getUserIdFromChannelName} from 'mattermost-redux/utils/channel_utils';
import {General} from 'mattermost-redux/constants';

import OverlayTrigger from 'components/overlay_trigger';
import Tooltip from 'components/tooltip';

import CallButton from 'plugins/call_button';
import ChannelHeaderPlug from 'plugins/channel_header_plug';
import {Constants, NotificationLevels} from 'utils/constants';
import {isEmptyObject} from 'utils/utils';
import type {GlobalState} from 'types/store';

import ChannelHeaderTitle from './channel_header_title';
import ChannelHeaderText from './channel_header_text';
import ChannelInfoButton from './channel_info_button';

const ChannelHeader = () => {
    const intl = useIntl();
    const dispatch = useDispatch();
    const channel = useSelector(getCurrentChannel) || {};
    const dmUser = useSelector((state: GlobalState) => {
        if (channel && channel.type === General.DM_CHANNEL) {
            const dmUserId = getUserIdFromChannelName(currentUser.id, channel.name);
            return getUser(state, dmUserId);
        }
        return undefined;
    })
    const doGetProfilesInChannel = makeGetProfilesInChannel();

    const gmMembers = useSelector((state: GlobalState) => {
        if (channel && channel.type === General.GM_CHANNEL) {
            return doGetProfilesInChannel(state, channel.id);
        }
        return undefined;
    });

    const currentUser = useSelector(getCurrentUser);
    const channelMember = useSelector(getMyCurrentChannelMembership);
    const isMuted = useSelector(isCurrentChannelMuted);

    useEffect(() => {
        dispatch(getCustomEmojisInText(channel ? channel.header : ''));
    }, [channel.header]);

    const unmute = useCallback(() => {
        if (!channelMember || !currentUser || !channel) {
            return;
        }

        const options = {mark_unread: NotificationLevels.ALL};
        dispatch(updateChannelNotifyProps(currentUser.id, channel.id, options));
    }, [channel, channelMember, currentUser]);

    const ariaLabelChannelHeader = intl.formatMessage({id:'accessibility.sections.channelHeader', defaultMessage: 'channel header region'});

    if (isEmptyObject(channel) ||
        isEmptyObject(channelMember) ||
        isEmptyObject(currentUser) ||
        (!dmUser && channel.type === Constants.DM_CHANNEL)
    ) {
        // Use an empty div to make sure the header's height stays constant
        return (
            <div className='channel-header'/>
        );
    }


    const channelMutedTooltip = (
        <Tooltip id='channelMutedTooltip'>
            <FormattedMessage
                id='channelHeader.unmute'
                defaultMessage='Unmute'
            />
        </Tooltip>
    );

    return (
        <div
            id='channel-header'
            aria-label={ariaLabelChannelHeader}
            role='banner'
            tabIndex={-1}
            data-channelid={`${channel.id}`}
            className='channel-header alt a11y__region'
            data-a11y-sort-order='8'
        >
            <div className='flex-parent'>
                <div className='flex-child'>
                    <div
                        id='channelHeaderInfo'
                        className='channel-header__info'
                    >
                        <div
                            className='channel-header__title dropdown'
                        >
                            <div>
                                <ChannelHeaderTitle
                                    dmUser={dmUser}
                                    gmMembers={gmMembers}
                                />
                            </div>
                            {isMuted &&
                                <OverlayTrigger
                                    delayShow={Constants.OVERLAY_TIME_DELAY}
                                    placement='bottom'
                                    overlay={channelMutedTooltip}
                                >
                                    <button
                                        id='toggleMute'
                                        onClick={unmute}
                                        className={'style--none color--link channel-header__mute inactive'}
                                        aria-label={intl.formatMessage({id: 'generic_icons.muted', defaultMessage: 'Muted Icon'})}
                                    >
                                        <i className={'icon icon-bell-off-outline'}/>
                                    </button>
                                </OverlayTrigger>}
                        </div>
                        <ChannelHeaderText dmUser={dmUser}/>
                    </div>
                </div>
                <ChannelHeaderPlug
                    channel={channel}
                    channelMember={channelMember}
                />
                <CallButton/>
                <ChannelInfoButton channel={channel}/>
            </div>
        </div>
    );
}

export default memo(ChannelHeader);
