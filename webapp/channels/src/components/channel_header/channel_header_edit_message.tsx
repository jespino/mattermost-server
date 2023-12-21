// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo, useCallback, useRef} from 'react';
import {useSelector, useDispatch} from 'react-redux';
import {useIntl} from 'react-intl';
import {Overlay} from 'react-bootstrap';
import {FormattedMessage} from 'react-intl';

import type {UserProfile} from '@mattermost/types/users';

import {Permissions} from 'mattermost-redux/constants';
import EditChannelHeaderModal from 'components/edit_channel_header_modal';
import {openModal} from 'actions/views/modals';

import ChannelPermissionGate from 'components/permissions_gates/channel_permission_gate';
import {getCurrentChannel} from 'mattermost-redux/selectors/entities/channels';
import {Constants, ModalIdentifiers} from 'utils/constants';
import {getCurrentTeamId} from 'mattermost-redux/selectors/entities/teams';

type Props = {
    dmUser?: UserProfile;
}

const ChannelHeaderEditMessage = ({
    dmUser,
}: Props) => {
    const teamId = useSelector(getCurrentTeamId)
    const intl = useIntl();
    const dispatch = useDispatch();
    const channel = useSelector(getCurrentChannel) || {};
    const isDirect = (channel.type === Constants.DM_CHANNEL);
    const isGroup = (channel.type === Constants.GM_CHANNEL);
    const isPrivate = (channel.type === Constants.PRIVATE_CHANNEL);
    const channelIsArchived = channel.delete_at !== 0;
    const headerOverlayRef = useRef<Overlay>(null);

    const showEditChannelHeaderModal = useCallback(() => {
        if (headerOverlayRef.current) {
            headerOverlayRef.current.hide();
        }

        const modalData = {
            modalId: ModalIdentifiers.EDIT_CHANNEL_HEADER,
            dialogType: EditChannelHeaderModal,
            dialogProps: {channel},
        };

        dispatch(openModal(modalData));
    }, [channel]);

    if (channelIsArchived) {
        return null
    }

    if (isDirect || isGroup) {
        if (isDirect && dmUser?.is_bot) {
            return null;
        }
        return (
            <button
                className='header-placeholder style--none'
                onClick={showEditChannelHeaderModal}
            >
                <FormattedMessage
                    id='channel_header.addChannelHeader'
                    defaultMessage='Add a channel header'
                />
                <i
                    className='icon icon-pencil-outline edit-icon'
                    aria-label={intl.formatMessage({id: 'channel_header.editLink', defaultMessage: 'Edit'})}
                />
            </button>
        );
    }
    return (
        <ChannelPermissionGate
            channelId={channel.id}
            teamId={teamId}
            permissions={[isPrivate ? Permissions.MANAGE_PRIVATE_CHANNEL_PROPERTIES : Permissions.MANAGE_PUBLIC_CHANNEL_PROPERTIES]}
        >
            <button
                className='header-placeholder style--none'
                onClick={showEditChannelHeaderModal}
            >
                <FormattedMessage
                    id='channel_header.addChannelHeader'
                    defaultMessage='Add a channel header'
                />
                <i
                    className='icon icon-pencil-outline edit-icon'
                    aria-label={intl.formatMessage({id: 'channel_header.editLink', defaultMessage: 'Edit'})}
                />
            </button>
        </ChannelPermissionGate>
    );
}

export default memo(ChannelHeaderEditMessage);
