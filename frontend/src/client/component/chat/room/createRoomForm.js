import React, {useState} from "react";
import axios from "axios";
import CardContent from "@material-ui/core/CardContent";
import TextField from "@material-ui/core/TextField";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import {makeStyles} from "@material-ui/core/styles";
import {FormControlLabel, Radio, RadioGroup} from "@material-ui/core";
import {Alert, AlertTitle} from '@material-ui/lab';

const useStyles = makeStyles((theme) => ({
    root: {
        width: 300,
        height: 450,
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

export function CreateRoomForm() {
    const [roomName, setRoomName] = useState('')
    const [info, setInfo] = useState('')
    const [pri, setPri] = useState('private')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')
    const [warning, setWarning] = useState('')
    const [success, setSuccess] = useState('')
    const classes = useStyles()

    async function handleSubmit() {
        if (roomName.length < 1) {
            setSuccess("")
            setWarning("please set room name.")
            return
        } else {
            setWarning("")
        }

        let data = {'room_name': roomName, 'info': info, 'private': pri, 'password': password}
        await axios.post('/api/createRoom', data)
            .then(res => {
                setSuccess("Created.")
                setError("")
                console.log("process")
            }).catch(error => {
                setSuccess("")
                setError("Unauthorized!")
            })
        console.log("Done")
        setRoomName('')
        setInfo('')
        setPassword('')
    }

    return (
        <div>
            <CardContent>
                {
                    error.length > 0 ?
                        <Alert severity="error">
                            <AlertTitle>Error</AlertTitle>
                            <strong>{error}</strong>
                        </Alert>
                        :
                        ""
                }
                {
                    warning.length > 0 ?
                        <Alert severity="warning">
                            <AlertTitle>Warning</AlertTitle>
                            <strong>{warning}</strong>
                        </Alert>
                        :
                        ""
                }
                {
                    success.length > 0 ?
                        <Alert severity="success">
                            <AlertTitle>Success</AlertTitle>
                            <strong>{success}</strong>
                        </Alert>
                        :
                        ""
                }


                <TextField error={warning.length > 0 && roomName.length < 1} value={roomName} className={classes.text}
                           required label="Room Name" name="room_name"
                           onChange={(event) => setRoomName(event.target.value)}/>
                <br/>
                <TextField value={info} className={classes.text} multiline
                           rows={4}
                           label="Information" type="text"
                           name="info" onChange={(event) => setInfo(event.target.value)}/>
                <RadioGroup value={pri} onChange={(event) => setPri(event.target.value)}>
                    <FormControlLabel value="private" control={<Radio/>} label="Private"/>
                    <FormControlLabel value="public" control={<Radio/>} label="Public"/>
                </RadioGroup>
                <TextField value={password} className={classes.text} disabled label="Password (Unavailable now)"
                           type="password"
                           name="password" onChange={(event) => setPassword(event.target.value)}/>

            </CardContent>
            <CardActions>
                <Button className={classes.button} variant="contained" color="primary" onClick={handleSubmit}>
                    Create Room
                </Button>
            </CardActions>

        </div>
    )

}