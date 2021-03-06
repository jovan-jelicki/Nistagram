import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import followersService from "../../services/followers.service";
import userService from "../../services/user.service";
import ProfileForSug from "../HomePage/ProfileForSug";
import {Button} from "react-bootstrap";
import Notification from "./Notification";


function Notifications() {
    const store = useSelector(state => state);

    const[notifications,setNotifications]=useState([])

    useEffect(() => {
        if(store.user.role === 'Admin' || store.user.role === "") window.location.replace("http://localhost:3000/unauthorized");
        getUserNotifications()
    }, []);

    async function getUserNotifications() {
        const response = await userService.getUserNotifications({
            id: '11fa801d-9064-4478-bdcd-0d522c32efa4',
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            setNotifications(response.data.notifications)
            console.log(response.data.notifications)
        } else {
            console.log("NIJE otpratio")
        }
    }

    return (
        <div style={{marginTop:'6%'}}>
            <Navigation/>
            <div style={{marginLeft:'20%', marginRight:'20%'}}>
                 <h3 style={{borderBottom:'1px solid black'}}>Notifications</h3>
                <div style={{marginTop:'4%'}}>
                    {notifications.map((notification, i) =>
                       <Notification creatorId={notification.creatorId} userId={notification.userId} text={notification.text} type={notification.type}/>
                    ) }
                </div>
            </div>
        </div>
    );
}

export default Notifications;