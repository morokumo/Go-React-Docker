import React, {useEffect} from "react";
import axios from "axios";
import Button from '@material-ui/core/Button';
import {Redirect} from "react-router-dom";

export function SignOut(props) {
    const [redirect, setRedirect] = React.useState(false);
    useEffect(() => {
    })

    function signOut() {
        let result = confirm("Sign out ?")
        if (result) {
            axios.post('api/signOut')
                .then((res) => {
                    console.log(res)
                    setRedirect(true)
                }).catch(() => {
                alert("Failed.")
            })
        }
    }

    return (
        <div>
            {redirect ? <Redirect to={'/'}/> : ''}
            <Button color="primary"  variant="contained" onClick={signOut}>Sign out</Button>
        </div>

    );
}