import React, {useState} from 'react';
import {Button, Form, Modal} from "react-bootstrap";
import ReCAPTCHA from "react-google-recaptcha";
import RegistrationPage from "./RegistrationPage";
import axios from "axios";
const TEST_SITE_KEY = "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI";


export function IndexPage(){
    const[details,setDetails]=useState({email:"", password:""});
    const[badCredentials,setCredentials]=useState(true);
    const[badEmail,setEmailErr]=useState(true);
    const[submitted,setSubmitted]=useState(false);
    const[reCaptcha,setCaptcha]=useState(0);
    const[logInDisabled,setLogInDisabled]=useState(false);
    const[showModal,setShowModal]=useState(false);

    const sendParams=()=>{
        axios
            .post("http://localhost:8001/users/login", {
                email: details.email,
                password: details.password
            })
            .then(res => {
                alert("Login successful ");
                //localStorage.setItem("jwt", res.data); // Change asap
            })
            .catch(res => {
                if (reCaptcha >= 2) {
                    setCaptcha(reCaptcha+1);
                    setLogInDisabled(true);
                    setCredentials(false);
                } else {
                    setCaptcha(reCaptcha+1);
                    setCredentials(false);
                }
            });
    }
    const submitHandler=e=>{
        setSubmitted(true);
        e.preventDefault();
        if(checkEmail(details.email)){
            setEmailErr(true);
            sendParams();
        }else{
            setEmailErr(false)
        }
    }

    const handleChange = event => {
        setDetails({
            ...details,
            [event.target.name]: event.target.value,
        });

        if(event.target.name==="email" && submitted){
            if(checkEmail(event.target.value)==false) setEmailErr(false);
            else setEmailErr(true);
        }
    };

    const checkEmail=(value)=>{
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }

    const closeCaptcha = () => {
        setCaptcha(0);
        setLogInDisabled(false);

    }

    const handleModal = () => {
        setShowModal(!showModal)
    }

    const closeModal = () => {
        setShowModal(!showModal)
    }


    return (
        <div style={{ padding: '60px 0', margin: '0 auto', maxWidth: '320px' }}>
            <br />
            <Form onSubmit={submitHandler}>
                <Form.Group size="lg" controlId="email">
                    <Form.Label>Email</Form.Label>
                    <Form.Control autoFocus type="email" name="email" onChange={handleChange}  value={details.email} />
                    <p hidden={badEmail} style={{ color: "red" }}> Email is not valid!</p>

                </Form.Group>
                <Form.Group size="lg" controlId="password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control  name="password" type="password" onChange={handleChange} value={details.password}/>
                </Form.Group>
                <p hidden={badCredentials} style={{ color: "red" }}> Invalid username or password!</p>
                {reCaptcha >= 3 &&
                <ReCAPTCHA
                    style={{ display: "inline-block" }}
                    theme="light"
                    ref={React.createRef()}
                    sitekey={TEST_SITE_KEY}
                    onChange={closeCaptcha}
                    // asyncScriptOnLoad={asyncScriptOnLoad}
                />
                }
                <Button disabled={logInDisabled} block size="lg" onClick={submitHandler}>Login </Button>
            </Form>
            <br />
            <div style={{ display: " table" }}>
                <p style={{ display: "table-cell" }}>Don't have account?</p>
                <a style={{ display: "table-cell" }} className="nav-link" style={{ 'color': '#00d8fe', 'fontWeight': 'bold' }} href='#' name="workHours" onClick={handleModal}>Register</a>
            </div>

            <Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Registration</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <RegistrationPage/>
                </Modal.Body>
            </Modal>
        </div>
    );
}