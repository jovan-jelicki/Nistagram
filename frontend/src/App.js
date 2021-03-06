import  React from "react";
import IndexPage from './pages/IndexPage.js'
import {ForgotPasswordPage} from './components/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import Chats from "./components/HomePageComponents/Chats";
import Saved from "./components/HomePageComponents/Saved";
import Search from "./pages/Search";
import Profile from "./components/ProfileComponent/Profile";
import NewPost from "./components/HomePageComponents/NewPost";
import Home from "./components/HomePage/Home";
import UnauthorizedPage from "./helpers/UnauthorizedPage";
import BlockedUsers from "./components/UserData/BlockedUsers";
import CloseFriends from "./components/UserData/CloseFriends";
import Liked from "./components/UserData/Liked";
import Disliked from "./components/UserData/Disliked";
import Archived from "./components/UserData/Archived";
import ProfileInfo from "./components/UserData/ProfileInfo";
import EditProfile from "./components/UserData/EditProfile";
import ChangePassword from "./components/UserData/ChangePassword";
import EditUserPrivacy from "./components/UserData/EditUserPrivacy";
import SubmitVerificationRequest from "./components/VerificationRequest/SubmitVerificationRequest";
import ViewMyVerificationRequests from "./components/VerificationRequest/ViewMyVerificationRequests";
import ViewPendingVerificationRequests from "./components/VerificationRequest/ViewPendingVerificationRequests";
import ViewAllVerificationRequests from "./components/VerificationRequest/ViewAllVerificationRequests";
import EditProfileImage from "./components/UserData/EditProfileImage";
import Notifications from "./components/Notifications/Notifications";


const App = () => {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/unauthorized' exact={true} component={UnauthorizedPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/home' exact  component={Home}/>
                <Route path='/search' exact={true} component={Search}/>
                <Route path='/profile/:username' exact component={Profile}/>
                <Route path='/info' exact component={ProfileInfo}/>

                <Route path='/newpost' exact component={NewPost} />
                <Route path='/chats' exact component={Chats} />
                <Route path='/saved' exact component={Saved} />
                <Route path='/notifications' exact component={Notifications} />
                <Route path='/submit-verification-request' exact component={SubmitVerificationRequest} />
                <Route path='/view-my-verification-request' exact component={ViewMyVerificationRequests} />
                <Route path='/view-pending-verification-request' exact component={ViewPendingVerificationRequests} />
                <Route path='/view-all-verification-request' exact component={ViewAllVerificationRequests} />


                <Route path='/blocked' exact component={BlockedUsers} />
                <Route path='/close_friends' exact component={CloseFriends} />
                <Route path='/liked' exact component={Liked} />
                <Route path='/disliked' exact component={Disliked} />
                <Route path='/archive' exact component={Archived} />
                <Route path='/edit_profile' exact component={EditProfile} />
                <Route path='/password' exact component={ChangePassword} />
                <Route path='/privacy' exact component={EditUserPrivacy} />
                <Route path='/edit_photo' exact component={EditProfileImage} />
            </Router>
        </div>
    );
}

export default App