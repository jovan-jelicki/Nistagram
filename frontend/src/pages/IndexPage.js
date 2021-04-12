import * as React from "react";
import {Button, Form, Modal,} from "react-bootstrap";

export default class IndexPage extends React.Component {
    constructor() {
        super();
        this.state = {
            password : '',
            email : '',
            user : !!localStorage.getItem("user") ? JSON.parse(localStorage.getItem("user")) : {},
            showModal : false,
            badCredentials : false
        }
    }

    render() {
        return (
            <div style={{padding : '60px 0', margin : '0 auto', maxWidth : '320px'}}>
                <br/>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Group size="lg" controlId="email">
                        <Form.Label>Email</Form.Label>
                        <Form.Control
                            autoFocus
                            type="email"
                            name="email"
                            value={this.state.email}
                            onChange={e => this.handleInputChange(e)}
                        />
                    </Form.Group>
                    <Form.Group size="lg" controlId="password">
                        <Form.Label>Password</Form.Label>
                        <Form.Control
                            name="password"
                            type="password"
                            value={this.state.password}
                            onChange={e => this.handleInputChange(e)}
                        />
                    </Form.Group>
                    <p hidden={this.state.badCredentials} style={{color : "red"}}> Invalid username or password!</p>
                    <div style={{display : "flex"}}>
                        <a href={'/#'} style={{float : "right"}}> Forgot password?</a>
                    </div>
                    <br/>
                    <br/>
                    <Button  block size="lg" disabled={!this.validateForm} onClick={this.handleSubmit}>
                        Login
                    </Button>
                </Form>
                <br/>
                <br/>
                <div style={{display : " table"}}>
                    <p style={{display: "table-cell"}}>Don't have account?</p>
                    <a style={{display: "table-cell"}} className="nav-link" style={{'color' : '#00d8fe', 'fontWeight' : 'bold'}} href='#' name="workHours" onClick={this.handleModal}>Register</a>
                </div>

                <Modal show={this.state.showModal} onHide={this.closeModal}  style={{'height':650}} >
                    <Modal.Header closeButton style={{'background':'silver'}}>
                        <Modal.Title>Registration</Modal.Title>
                    </Modal.Header>
                    <Modal.Body style={{'background':'silver'}}>
                    </Modal.Body>
                </Modal>


            </div>
        )
    }

    validateForm = () => {
        return this.state.email.length > 0 && this.state.password.length > 0;
    }

    handleModal=()=>{
        this.setState({
            showModal: !this.state.showModal,
        });
    }

    closeModal=()=>{
        this.setState({
            showModal : !this.state.showModal
        });
    }


    handleInputChange = (event) => {
        const target = event.target;
        this.setState({
            [target.name] : target.value,
        })
    }

    handleSubmit = () => {

    }

}