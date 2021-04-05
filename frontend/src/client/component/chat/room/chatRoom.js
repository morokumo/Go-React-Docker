import React, {useEffect, useState} from "react";
import axios from "axios";
import 'regenerator-runtime/runtime';

import {
    Button,
    Card,
    CardActions,
    CardContent,
    Divider,
    ListItem,
    ListItemText,
    makeStyles,
    TextField,
    Typography
} from "@material-ui/core";
import SendIcon from '@material-ui/icons/Send';
import {FixedSizeList} from "react-window";

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 800,
        minWidth: 300,
        height: 700,
        backgroundColor: theme.palette.background.paper,
    },
    content: {
        width: '100%',
        maxWidth: 780,
        minWidth: 320,
        maxHeight: 600,
        minHeight: 700,
        backgroundColor: theme.palette.background.paper,
    },
    prime: {
        fontSize: 14
    },
    date: {
        fontSize: 12
    },
    second: {
        fontSize: 18
    }
}));

export function ChatRoom(props) {
    const [room, setRoom] = useState(props.room);
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);
    const [list, setList] = useState();
    const [error, setError] = useState();

    const classes = useStyles()
    useEffect(() => {
        props.socket.on("get", (e) => {
            if (e['error'] === undefined) {
                setMessages(old => [...old, e])
            } else {
                setError(e['error'])
            }
            console.log(messages, e)

            console.log("SEND")
        }, [])
    }, [])

    useEffect(() => {
        getMessage()
        setMessage('')
        props.socket.emit("join", room.ID)
        console.log("JOIN", props.socket.id)
    }, [room]);

    useEffect(() => {
        if (list !== undefined) {
            list.scrollToItem(messages.length, 'end')
        }
    }, [messages.length])


    useEffect(() => {
        setRoom(props.room)
    }, [props.room]);


    async function handleSubmit() {
        if (message.length > 0) {
            // await sendMessage()
            // await getMessage()
            props.socket.emit("send", {room: room.ID, id: props.socket.id, message: message})
            list.scrollToItem(messages.length, 'end')
        }
        setMessage('')
    }

    async function getMessage() {
        await axios.post('/api/getMessage', {'room_id': room.ID, 'message': message})
            .then(res => {
                console.log(res.data)
                if (!(res.data === null) && !(res.data === undefined)) {
                    setMessages(res.data)
                }
            }).catch(error => {

            })
    }

    const Row = ({index, style}) => (
        <div>
            {/*<ListSubheader></ListSubheader>*/}
            <ListItem style={style}>
                <ListItemText
                    primary={
                        <div className={classes.prime}>
                            <b>{messages[index].AccountID}</b>
                            {'　　　'}
                            {messages[index].SendTime}
                        </div>
                    }
                    secondary={
                        <Typography fontFamily="Monospace" className={classes.second}>
                            {messages[index].Text}
                        </Typography>
                    }/>
            </ListItem>
        </div>
    );

    return (
        <div>
            <Card className={classes.root}>
                <CardContent>
                    <h1>{room.Name}</h1>
                    <Divider/>
                    <span>{room.Info}</span>
                    <Divider/>

                    <FixedSizeList height={400} width={600} itemSize={60} itemCount={messages.length}
                                   ref={(fl) => {
                                       setList(fl)
                                       // fl.scrollToItem(messages.length, 'end')
                                   }}
                    >
                        {Row}
                    </FixedSizeList>
                </CardContent>
                <Divider/>
                <CardActions>
                    {error}
                    <TextField value={message} required id="standard-required" label="Message"
                               onChange={(event) => setMessage(event.target.value)}
                               onKeyPress={(i) => i.key === 'Enter' ? handleSubmit() : ''}
                    />
                    <Button onClick={handleSubmit}>
                        <SendIcon/>
                    </Button>
                </CardActions>
            </Card>


        </div>
    );
}