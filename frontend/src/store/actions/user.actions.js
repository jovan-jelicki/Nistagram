import { userConstants } from './../constants'

const registerRequest = (data) => {
    return { type: userConstants.REGISTER_REQUEST, data }
}

const loginRequest = (data) => {
    return { type: userConstants.LOGIN_REQUEST, data }
}

const followRequest = (data) => {
    return { type: userConstants.FOLLOW_REQUEST, data }
}

export const userActions = {
    registerRequest,
    loginRequest,
    followRequest
}