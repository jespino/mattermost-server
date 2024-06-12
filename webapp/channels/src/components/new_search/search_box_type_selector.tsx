// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {FormattedMessage} from 'react-intl';
import {useSelector} from 'react-redux';
import styled from 'styled-components';

import type {GlobalState} from 'types/store';

const SearchTypeSelectorContainer = styled.div`
    margin: 24px 32px 0px 24px;
    display: flex;
    align-items: center;
    padding: 4px;
    background-color: var(--center-channel-bg);
    border-radius: var(--radius-m);
    border: var(--border-default);
    width: fit-content;
    gap: 4px;
`;

type SearchTypeItemProps = {
    selected: boolean;
};

const SearchTypeItem = styled.button<SearchTypeItemProps>`
    display: flex;
    cursor: pointer;
    padding: 4px 10px;
    background-color: ${(props) => (props.selected ? 'rgba(var(--button-bg-rgb), 0.08)' : 'transparent')};
    color: ${(props) => (props.selected ? 'var(--button-bg)' : 'rgba(var(--center-channel-color-rgb), 0.75)')};
    border-radius: 4px;
    font-size: 12px;
    line-height: 16px;
    font-weight: 600;
    border: none;
    &:hover {
        color: rgba(var(--center-channel-color-rgb), 0.88);
        background: rgba(var(--center-channel-color-rgb), 0.08);
    }
`;

type Props = {
    searchType: string;
    setSearchType: (searchType: string) => void;
}

const SearchTypeSelector = ({searchType, setSearchType}: Props) => {
    const SearchPluginButtons = useSelector((state: GlobalState) => state.plugins.components.SearchButtons) || [];
    return (
        <SearchTypeSelectorContainer>
            <SearchTypeItem
                selected={searchType === 'messages'}
                onClick={() => setSearchType('messages')}
            >
                <FormattedMessage
                    id='search_bar.usage.search_type_messages'
                    defaultMessage='Messages'
                />
            </SearchTypeItem>
            <SearchTypeItem
                selected={searchType === 'files'}
                onClick={() => setSearchType('files')}
            >
                <FormattedMessage
                    id='search_bar.usage.search_type_files'
                    defaultMessage='Files'
                />
            </SearchTypeItem>
            <SearchTypeItem
                selected={searchType === 'omnisearch'}
                onClick={() => setSearchType('omnisearch')}
            >
                <FormattedMessage
                    id='search_bar.usage.search_type_omnisearch'
                    defaultMessage='OmniSearch'
                />
            </SearchTypeItem>
            {SearchPluginButtons.map(({component, pluginId}: any) => {
                const Component = component as React.ComponentType;
                return (
                    <SearchTypeItem
                        key={pluginId}
                        selected={searchType === pluginId}
                        onClick={() => setSearchType(pluginId)}
                    >
                        <Component/>
                    </SearchTypeItem>
                );
            })}
        </SearchTypeSelectorContainer>
    );
};

export default SearchTypeSelector;
