// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import styled from 'styled-components';
import Markdown from 'components/markdown';

type Props = {
    icon: string
    link: string
    title: string
    subtitle: string
    description: string
}

const OmniSearchResultItemContainer = styled.div`
    display: flex;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid var(--center-channel-color-08);
`

const Icon = styled.img`
    width: 40px;
    height: 40px;
    margin-right: 10px;
    align-self: flex-start;
    margin-top: 5px;
    border-radius: 20px;
`

const Title = styled.div`
    font-weight: 600;
    font-size: 16px;
    margin-bottom: 5px;
    a {
        color: var(--center-channel-color);
        text-decoration: none;
        cursor: pointer;
    }
`

const Subtitle = styled.div`
    font-weight: 400;
    font-size: 14px;
    margin-bottom: 5px;
`

const Description= styled.div`
    max-height: 100px;
    overflow: hidden;
`

const OmniSearchResultItem = ({icon, link, title, subtitle, description}: Props) => {
    return (
        <OmniSearchResultItemContainer>
            <Icon src={icon}/>
            <div>
                <Title><a href={link} target="_blank">{title}</a></Title>
                {subtitle && <Subtitle>{subtitle}</Subtitle>}
                <Description>
                    <Markdown message={description}/>
                </Description>
            </div>
        </OmniSearchResultItemContainer>
    )
}

export default OmniSearchResultItem;
