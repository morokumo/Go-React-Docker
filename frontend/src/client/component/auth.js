import React, {useEffect, useState} from 'react'
import axios from "axios";
import {Redirect} from "react-router-dom";

export default function Auth(props) {
    const [auth, setAuth] = useState(true)

    useEffect(() => {
        verify().then((ok) => {
            setAuth(ok)
        }).catch((ok) => {
            setAuth(false)
        })
    }, [props])

    async function verify() {
        let ok = false
        await axios.post('api/auth')
            .then((response) => {
                if (response.data.message === 'Status OK') {
                    ok = true;
                }
            })
        return ok
    }

    return (
        <div>
            {auth ? '' : <Redirect to={'/'}/>}
            <div>
                {props.children}
            </div>

        </div>
    );
}