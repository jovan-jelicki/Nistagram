import {Button, FormControl, Table} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import axios from "axios";

const EditProfile = props => {
    const [user, setUser] = useState(props.user);  //  PAZI NA KUPLJENJE PROPS-A DOLAZI MU user{username:''}
    const[edit,setEdit]=useState(false)


    function editProfile(){ //501
        axios
            .post('http://localhost:8080/api/users/api/users/update_profile', {
        user:{
            id: '11e215ef-a6ce-40c6-ad0b-e602e298f0a4',
            firstName: 'Ana',
            lastName: 'Zoric',
            email: 'kkosta98@gmail.com',
            phoneNumber: '0658456152',
            username: 'kole998',
            profilePhoto: 'idk',
            sex: 'MAN',
            website:'igorijovanzauvekzajednojej',
            isActive: true
        }
            })
            .then(res => {
                console.log("RADI")

            }).catch(res => {
            console.log("NE RADIs")
        })
    }
    function handleInputChange(event) {
        setUser({
            ...user,
            [event.target.name]: event.target.value,
        });
    }
    function  activateUpdateMode(){
        setEdit(true);
    }

        return(
            <div >
            <Table style={{marginLeft:'850px', display:'inline-block', float:'right',alignContent:'', maxWidth: '900px'}}>
                <tbody>
                <tr>
                    <td>Username</td>
                    {edit
                        ? <FormControl name="username" className="mt-2 mb-1" value={user.username} onChange={handleInputChange}/>
                        : <td>{user.username}</td>
                    }
                </tr>
                <tr>
                    <td>First Name</td>
                    {edit
                        ? <FormControl name="firstName" className="mt-2 mb-1" value={user.firstName} onChange={handleInputChange}/>
                        : <td>{user.firstName}</td>
                    }
                </tr>
                <tr>
                    <td>Last Name</td>
                    {edit
                        ? <FormControl name="lastName" className="mt-2 mb-2" value={user.lastName} onChange={handleInputChange}/>
                        : <td>{user.lastName}</td>
                    }
                </tr>
                <tr>
                    <td>Email</td>
                    <td>{user.email}</td>
                </tr>
                <tr>
                    <td>Birth date</td>
                    {edit
                        ? <FormControl name="birthDate" className="mt-2 mb-2" value={user.birthDate} onChange={handleInputChange}/>
                        : <td>{user.birthDate}</td>
                    }
                </tr>
                <tr>
                    <td>Phone Number</td>
                    {edit
                        ? <FormControl name="phoneNumber" className="mt-2 mb-2" value={user.phoneNumber} onChange={handleInputChange}/>
                        : <td>{user.phoneNumber}</td>
                    }
                </tr>
                <tr>
                    <td>Sex</td>
                    {edit
                        ? <FormControl name="sex" className="mt-2 mb-2" value={user.sex} onChange={handleInputChange}/>
                        : <td>{user.sex}</td>
                    }
                </tr>
                <tr>
                    <td>Biography</td>
                    {edit
                        ? <FormControl name="biography" className="mt-2 mb-2" value={user.biography} onChange={handleInputChange}/>
                        : <td>{user.biography}</td>
                    }
                </tr>
                <tr>
                    <td>Web site</td>
                    {edit
                        ? <FormControl name="website" className="mt-2 mb-2" value={user.website} onChange={handleInputChange}/>
                        : <td>{user.website}</td>
                    }
                </tr>
                {edit &&
                <tr >
                    <button style={{margin:'10px'}} type="button" className="btn btn-primary btn-sm" onClick={editProfile}>Save</button>
                    <button style={{marginLeft:'10px'}} type="button"  className="btn btn-primary btn-sm" onClick={() =>  setEdit(false)}>Cancel
                    </button>
                </tr>
                }
                </tbody>
            </Table>
                {!edit &&
                    <button  style={{marginRight:'100px', float:'right'}} type="button" className="btn btn-outline-danger" onClick={activateUpdateMode}>Edit profile</button>
                }
            </div>
        );
};export default EditProfile;