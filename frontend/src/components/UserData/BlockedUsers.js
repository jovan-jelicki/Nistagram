import React, { useState } from 'react';
import ProfileInfo from "./ProfileInfo";

function BlockedUsers() {
    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>

            </div>
        </div>
    );
}

export default BlockedUsers;