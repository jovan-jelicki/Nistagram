import React from "react";
import {Button} from "react-bootstrap";
//ovde dolazim ako sam usla u profil nekog drugog usera
//treba da mu posaljem moj id, i da uzmem njegov id kako bih proverila da li se pratimo, ili ne
function FollowAndUnfollow(){
    return(
        <div>
            <Button variant="primary" style={{margin: "10px"}}>Follow</Button>
            <Button style={{margin: "10px"}}>UnFollow</Button>
        </div>
    );
}export default FollowAndUnfollow;