import React, {useEffect, useState} from "react";
import axios from "axios";
import {Redirect} from "react-router-dom";
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import {makeStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme) => ({
    root: {
        width: 300,
        height: 350,
        margin: "0 auto"
    },
    bullet: {
        display: 'inline-block',
        margin: '0 2px',
        transform: 'scale(0.8)',
    },
    title: {
        fontSize: 25,
    },
    pos: {
        // marginBottom: 12,
        // marginLeft:"auto",
        // marginRight:"auto"

    },
    button: {
        margin: theme.spacing(1),
    },
    text: {
        margin: theme.spacing(1),
        width: '25ch',
    }
}));


export function SignForm(props) {
    const [accountID, setAccountID] = useState('')
    const [password, setPassword] = useState('')
    const [auth, setAuth] = useState(false)
    const classes = useStyles();

    useEffect(() => {
    })

    function getData() {
        return {'account_id': accountID, 'password': password}
    }

    function onSubmitSignUP() {
        let data = getData()
        if (data['account_id'] === '' || data['password'] === '') {
            alert("please input")
        } else {
            axios.post('api/signUp', data)
                .then(() => {
                    setAuth(true)
                }).catch(() => {
                alert("already exist.")
            })
        }

    }

    function onSubmitSignIN() {
        let data = getData()
        console.log('AUTH')
        if (data['account_id'] === '' || data['password'] === '') {
            alert("please input")
        } else {
            axios.post('api/signIn', data)
                .then(() => {
                    setAuth(true)
                }).catch(() => {
                alert("account incorrect")
            })
        }
    }

    return (
        <div>
            {auth ? <Redirect to={'/'}/> : ''}
            <Card className={classes.root}>
                <CardContent>
                    <Typography className={classes.title} color="textSecondary" gutterBottom>
                        Sign Form
                    </Typography>
                    <TextField className={classes.text} required id="standard-required" label="Account ID"
                               name="account_id" variant="filled"
                               onChange={(event) => setAccountID(event.target.value)}/>

                    <br/>
                    <TextField className={classes.text} required id="standard-required" label="Password"
                               name="account_id" variant="filled"
                               onChange={(event) => setPassword(event.target.value)}/>
                </CardContent>
                <CardActions>
                    <Button className={classes.button} variant="contained" color="primary" onClick={onSubmitSignIN}>
                        Sign in
                    </Button>
                    <Button className={classes.button} variant="contained" color="secondary" onClick={onSubmitSignUP}>
                        Sign up
                    </Button>
                </CardActions>
            </Card>
        </div>

    );
}