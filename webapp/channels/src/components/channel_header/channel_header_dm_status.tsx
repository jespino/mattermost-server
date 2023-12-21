// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo} from 'react';
import {useSelector} from 'react-redux';
import {FormattedMessage} from 'react-intl';

import type {UserProfile} from '@mattermost/types/users';

import {displayLastActiveLabel, getLastActivityForUserId, getLastActiveTimestampUnits} from 'mattermost-redux/selectors/entities/users';
import {makeGetCustomStatus} from 'selectors/views/custom_status';

import {toTitleCase} from 'utils/utils';

import Timestamp from 'components/timestamp';
import StatusIcon from 'components/status_icon';
import {getCurrentChannel} from 'mattermost-redux/selectors/entities/channels';
import {Constants} from 'utils/constants';
import type {GlobalState} from 'types/store';

import ChannelHeaderCustomStatus from './channel_header_custom_status';

type Props = {
    dmUser?: UserProfile;
}

const ChannelHeaderDmStatus = ({
    dmUser,
}: Props) => {
    const channel = useSelector(getCurrentChannel) || {};
    const isDirect = (channel.type === Constants.DM_CHANNEL);
    const isLastActiveEnabled = useSelector((state: GlobalState) => dmUser ? displayLastActiveLabel(state, dmUser.id) : false);
    const lastActivityTimestamp = useSelector((state: GlobalState) => dmUser ? getLastActivityForUserId(state, dmUser.id) : undefined);
    const timestampUnits = useSelector((state: GlobalState) => dmUser ? getLastActiveTimestampUnits(state, dmUser.id) : undefined);
    const getCustomStatus = makeGetCustomStatus();
    const customStatus = useSelector((state: GlobalState) => dmUser ? getCustomStatus(state, dmUser.id) : undefined)

    if (!isDirect || dmUser?.delete_at || dmUser?.is_bot) {
        return null;
    }

    if (isLastActiveEnabled && lastActivityTimestamp && timestampUnits) {
        return (
            <>
                <StatusIcon status={channel.status}/>
                <span className='header-status__text'>
                    <span className='last-active__text'>
                        <FormattedMessage
                            id='channel_header.lastActive'
                            defaultMessage='Active {timestamp}'
                            values={{
                                timestamp: (
                                    <Timestamp
                                        value={lastActivityTimestamp}
                                        units={timestampUnits}
                                        useTime={false}
                                        style={'short'}
                                    />
                                ),
                            }}
                        />
                    </span>
                    <ChannelHeaderCustomStatus
                        dmUser={dmUser}
                        customStatus={customStatus}
                    />
                </span>
            </>
        )
    }
    return (
        <>
            <StatusIcon status={channel.status}/>
            <span className='header-status__text'>
                <FormattedMessage
                    id={`status_dropdown.set_${channel.status}`}
                    defaultMessage={toTitleCase(channel.status || '')}
                />
                <ChannelHeaderCustomStatus
                    dmUser={dmUser}
                    customStatus={customStatus}
                />
            </span>
        </>
    );
}

export default memo(ChannelHeaderDmStatus);
