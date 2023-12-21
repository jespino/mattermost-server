// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo} from 'react';
import {useSelector} from 'react-redux';

import type {UserCustomStatus, UserProfile} from '@mattermost/types/users';

import {isCustomStatusEnabled, isCustomStatusExpired} from 'selectors/views/custom_status';

import CustomStatusEmoji from 'components/custom_status/custom_status_emoji';
import CustomStatusText from 'components/custom_status/custom_status_text';

type Props = {
    dmUser?: UserProfile;
    customStatus?: UserCustomStatus;
}

const ChannelHeaderCustomStatus = ({
    dmUser,
    customStatus,
}: Props) => {
    const isCustomStatusEnabledValue = useSelector(isCustomStatusEnabled);
    const isCustomStatusExpiredValue = useSelector(isCustomStatusExpired(customStatus));

    const isStatusSet = !isCustomStatusExpiredValue && (customStatus?.text || customStatus?.emoji);
    if (!(isCustomStatusEnabledValue && isStatusSet)) {
        return null;
    }

    return (
        <div className='custom-emoji__wrapper'>
            <CustomStatusEmoji
                userID={dmUser?.id}
                showTooltip={true}
                tooltipDirection='bottom'
                emojiStyle={{
                    verticalAlign: 'top',
                    margin: '0 4px 1px',
                }}
            />
            <CustomStatusText
                text={customStatus?.text}
            />
        </div>
    );
};

export default memo(ChannelHeaderCustomStatus);

