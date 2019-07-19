import React from 'react';

import Avatar from '../Avatar/Avatar';
import { Link } from 'react-router-dom';

import "./PostHeader.css";
import UserInfoPopup from '../UserInfoPopup/UserInfoPopup';

export default function PostHeader({ userID, fullname, username, createdAt }) {
    return (
        <div className="PostHeader">
            <div className="PostHeader__Avatar">
                <Avatar width="2.75rem" userID={userID} />
            </div>

            <div className="PostHeader__Info">
                <UserInfoPopup username={username} trigger={<div className="PostHeader__Info__Fullname"><Link to={`/${username}`}>{fullname}</Link></div>} />
                <div className="PostHeader__Info__CreatedAt">{createdAt}</div>
            </div>
        </div>
    )
}
