// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import styled from 'styled-components';

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
`

const OmniSearchResultItem = ({icon, link, title, subtitle, description}: Props) => {
    return (
        <OmniSearchResultItemContainer>
            <img src={icon}/>
            <div>
                <h1><a href={link} target="_blank">{title}</a></h1>
                {subtitle && <h2>{subtitle}</h2>}
                <p>{description}</p>
            </div>
        </OmniSearchResultItemContainer>
    )
}

export default OmniSearchResultItem;
