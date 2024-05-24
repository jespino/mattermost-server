// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {DateTime} from 'luxon';
import React from 'react';
import type {ReactNode} from 'react';
import {FormattedMessage, useIntl} from 'react-intl';
import styled from 'styled-components';

import type {UserProfile} from '@mattermost/types/users';

import {isGuest} from 'mattermost-redux/utils/user_utils';
import {getTimezoneForUserProfile} from 'mattermost-redux/selectors/entities/timezone';

import CustomStatusEmoji from 'components/custom_status/custom_status_emoji';
import SharedUserIndicator from 'components/shared_user_indicator';
import StatusIcon from 'components/status_icon';
import BotTag from 'components/widgets/tag/bot_tag';
import GuestTag from 'components/widgets/tag/guest_tag';
import Tag from 'components/widgets/tag/tag';
import Avatar from 'components/widgets/users/avatar';
import Timestamp from 'components/timestamp';
import Moon from 'components/common/svg_images_components/moon_svg';
import WithTooltip from 'components/with_tooltip';

import {Constants} from 'utils/constants';
import * as Utils from 'utils/utils';

import {SuggestionContainer} from '../suggestion';
import type {SuggestionProps} from '../suggestion';

export interface Item extends UserProfile {
    display_name: string;
    name: string;
    isCurrentUser: boolean;
    type: string;
}

interface Group extends Item {
    member_count: number;
}

const MoonIcon = styled(Moon)`
    svg {
        width: 20px;
        height: 20px;
    }
    svg path {
        fill: rgba(var(--center-channel-color-rgb), 0.76);
    }
`

const AtMentionSuggestion = React.forwardRef<HTMLDivElement, SuggestionProps<Item>>((props, ref) => {
    const {item} = props;

    const intl = useIntl();

    let itemname: string;
    let localTime: ReactNode = null;
    let description: ReactNode;
    let icon: JSX.Element;
    let customStatus: ReactNode;
    if (item.username === 'all') {
        itemname = 'all';
        description = (
            <FormattedMessage
                id='suggestion.mention.all'
                defaultMessage='Notifies everyone in this channel'
            />
        );
        icon = (
            <span className='suggestion-list__icon suggestion-list__icon--large'>
                <i
                    className='icon icon-account-multiple-outline'
                    title={intl.formatMessage({id: 'generic_icons.member', defaultMessage: 'Member Icon'})}
                />
            </span>
        );
    } else if (item.username === 'channel') {
        itemname = 'channel';
        description = (
            <FormattedMessage
                id='suggestion.mention.channel'
                defaultMessage='Notifies everyone in this channel'
            />
        );
        icon = (
            <span className='suggestion-list__icon suggestion-list__icon--large'>
                <i
                    className='icon icon-account-multiple-outline'
                    title={intl.formatMessage({id: 'generic_icons.member', defaultMessage: 'Member Icon'})}
                />
            </span>
        );
    } else if (item.username === 'here') {
        itemname = 'here';
        description = (
            <FormattedMessage
                id='suggestion.mention.here'
                defaultMessage='Notifies everyone online in this channel'
            />
        );
        icon = (
            <span className='suggestion-list__icon suggestion-list__icon--large'>
                <i
                    className='icon icon-account-multiple-outline'
                    title={intl.formatMessage({id: 'generic_icons.member', defaultMessage: 'Member Icon'})}
                />
            </span>
        );
    } else if (item.type === Constants.MENTION_GROUPS) {
        itemname = item.name;
        description = (
            <span className='ml-1'>{'- '}{item.display_name}</span>
        );
        icon = (
            <span className='suggestion-list__icon suggestion-list__icon--large'>
                <i
                    className='icon icon-account-multiple-outline'
                    title={intl.formatMessage({id: 'generic_icons.member', defaultMessage: 'Member Icon'})}
                />
            </span>
        );
    } else {
        itemname = item.username;

        if (item.isCurrentUser) {
            if (item.first_name || item.last_name) {
                description = <span>{Utils.getFullName(item)}</span>;
            }
        } else if (item.first_name || item.last_name || item.nickname) {
            description = <span>{`${Utils.getFullName(item)} ${item.nickname ? `(${item.nickname})` : ''}`.trim()}</span>;
        }

        icon = (
            <span className='status-wrapper style--none'>
                <span className='profile-icon'>
                    <Avatar
                        username={item && item.username}
                        size='sm'
                        url={Utils.imageURLForUser(item.id, item.last_picture_update)}
                    />
                </span>
                <StatusIcon status={item && item.status}/>
            </span>
        );

        const itemTimezone = getTimezoneForUserProfile(item);
        const itemUserDate = DateTime.local().setZone(itemTimezone.useAutomaticTimezone ? itemTimezone.automaticTimezone : itemTimezone.manualTimezone);

        if (!item.isCurrentUser && (itemUserDate.get('hour') >= Constants.REMOTE_USERS_HOUR_LIMIT_END_OF_THE_DAY || itemUserDate.get('hour') < Constants.REMOTE_USERS_HOUR_LIMIT_BEGINNING_OF_THE_DAY)) {
            localTime = (
                <div>
                    <WithTooltip
                        id={'localTimeTooltip'}
                        title={
                            <Timestamp
                                useRelative={false}
                                value={itemUserDate.toMillis()}
                                useDate={false}
                                userTimezone={itemTimezone}
                                useTime={{
                                    hour: 'numeric',
                                    minute: 'numeric',
                                }}
                            />
                        }
                        placement='right'
                    >
                        <div>
                            <MoonIcon/>
                        </div>
                    </WithTooltip>
                </div>
            )
        }

        customStatus = (
            <CustomStatusEmoji
                showTooltip={true}
                userID={item.id}
                emojiSize={15}
                emojiStyle={{
                    margin: '0 4px 4px',
                }}
            />
        );
    }

    const youElement = item.isCurrentUser ? (
        <FormattedMessage
            id='suggestion.user.isCurrent'
            defaultMessage='(you)'
        />
    ) : null;

    const sharedIcon = item.remote_id ? (
        <SharedUserIndicator
            id={`sharedUserIndicator-${item.id}`}
            className='shared-user-icon'
        />
    ) : null;

    let countBadge;
    if (item.type === Constants.MENTION_GROUPS) {
        countBadge = (
            <span className='suggestion-list__group-count'>
                <Tag
                    text={
                        <FormattedMessage
                            id='suggestion.group.members'
                            defaultMessage='{member_count} {member_count, plural, one {member} other {members}}'
                            values={{
                                member_count: (item as Group).member_count,
                            }}
                        />
                    }
                />
            </span>
        );
    }

    return (
        <SuggestionContainer
            ref={ref}
            {...props}
            data-testid={`mentionSuggestion_${itemname}`}
        >
            {icon}
            <span className='suggestion-list__ellipsis'>
                <span className='suggestion-list__main'>
                    {'@' + itemname}
                </span>
                {item.is_bot && <BotTag/>}
                {description}
                {youElement}
                {customStatus}
                {sharedIcon}
                {isGuest(item.roles) && <GuestTag/>}
                {localTime}
            </span>
            {countBadge}
        </SuggestionContainer>
    );
});

AtMentionSuggestion.displayName = 'AtMentionSuggestion';
export default AtMentionSuggestion;
