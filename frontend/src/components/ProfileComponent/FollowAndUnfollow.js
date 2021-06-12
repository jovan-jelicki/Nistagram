import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import axios from "axios";
//ovde dolazim ako sam usla u profil nekog drugog usera
//treba da mu posaljem moj id, i da uzmem njegov id kako bih proverila da li se pratimo, ili ne
function FollowAndUnfollow(props){
    const{user,loggedUser,followers,getFollowers}=props;
    const[follows,setFollows]=useState(false);
    console.log(loggedUser)

    useEffect(() => {
        setFollows(followers.some(item=>item.UserId===loggedUser.id))
        console.log(follows)
    }, [followers])

    function follow(){
        console.log({
            userId :loggedUser.id,
            followerId : user.id,
            isMuted :false,
            isCloseFriends :false,
            isApprovedRequest :true,
            isNotificationEnabled : true
         })

        axios
            .post('http://localhost:8005/api/followers/create_connection', {
                 follower: {
                    userId :loggedUser.id,
                    followerId : user.id,
                    isMuted :false,
                    isCloseFriends :false,
                    isApprovedRequest :true,
                    isNotificationEnabled : true
                 }
            })
            .then(res => {
              console.log("ZAPRACENO")
                props.getFollowers();
            }).catch(res => {
            console.log("NIJE ZAPRACENO")
        })
    }

    function unfollow(){
        axios
            .post('http://localhost:8005/api/followers/delete_directed', {
                follower: {
                    UserId :loggedUser.id,
                    FollowerId : user.id,
                    IsMuted :false,
                    IsCloseFriends :false,
                    IsApprovedRequest :true,
                    IsNotificationEnabled : true
                }
            })
            .then(res => {
                console.log("otpratio")
                props.getFollowers();
            }).catch(res => {
            console.log("NIJE otpratio")
        })
    }
    return(
        <div>
            {!follows ?
                <Button variant="primary" style={{margin: "10px"}} onClick={follow}>Follow</Button>
                :
                <Button style={{margin: "10px"}} onClick={unfollow}>UnFollow</Button>
            }
        </div>
    );
}export default FollowAndUnfollow;