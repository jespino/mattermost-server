// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React, {memo, useCallback} from 'react';
import {useSelector, useDispatch} from 'react-redux';

import {getCurrentChannelStats, getCurrentChannel} from 'mattermost-redux/selectors/entities/channels';

import {closeRightHandSide, showChannelMembers} from 'actions/views/rhs';
import {getRhsState} from 'selectors/rhs';

import {Constants, RHSStates} from 'utils/constants';

import HeaderIconWrapper from './components/header_icon_wrapper';

const EMPTY_CHANNEL_STATS = {member_count: 0, guest_count: 0, pinnedpost_count: 0, files_count: 0};

const ChannelHeaderMembersButton = () => {
    const dispatch = useDispatch();
    const channel = useSelector(getCurrentChannel) || {};
    const rhsState = useSelector(getRhsState);
    const stats = useSelector(getCurrentChannelStats) || EMPTY_CHANNEL_STATS;
    const memberCount = stats.member_count;
    const isDirect = (channel.type === Constants.DM_CHANNEL);

    const toggleChannelMembersRHS = useCallback(() => {
        if (rhsState === RHSStates.CHANNEL_MEMBERS) {
            dispatch(closeRightHandSide());
        } else {
            dispatch(showChannelMembers(channel.id));
        }
    }, [channel.id]);

    if (isDirect) {
        return null;
    }
    const membersIconClass = classNames('member-rhs__trigger channel-header__icon channel-header__icon--left channel-header__icon--wide', {
        'channel-header__icon--active': rhsState === RHSStates.CHANNEL_MEMBERS,
    });

    const membersIcon = (
        <>
            <i
                aria-hidden='true'
                className='icon icon-account-outline channel-header__members'
            />
            <span
                id='channelMemberCountText'
                className='icon__text'
            >
                {memberCount || '-'}
            </span>
        </>
    );

    return (
        <HeaderIconWrapper
            iconComponent={membersIcon}
            ariaLabel={true}
            buttonClass={membersIconClass}
            buttonId={'member_rhs'}
            onClick={toggleChannelMembersRHS}
            tooltipKey={'channelMembers'}
        />
    );
};

export default memo(ChannelHeaderMembersButton);
