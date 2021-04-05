import React from "react";
import {
    Button,
    Card,
    Collapse,
    Divider,
    List,
    ListItem,
    ListItemIcon,
    ListItemText,
    makeStyles
} from "@material-ui/core";
import {FindMyRooms} from "./room/findMyRooms";
import {CreateRoomForm} from "./room/createRoomForm";
import {ExpandLess, ExpandMore} from "@material-ui/icons";
import AddBoxIcon from '@material-ui/icons/AddBox';
import {FindPublicRooms} from "./room/findPublicRooms";
import axios from "axios";
import {SignOut} from "../signOut";

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 400,
        minWidth: 350,
        height: 700,
        backgroundColor: theme.palette.background.paper,
    },
}));
export default function LeftNav(props) {
    const classes = useStyles()
    const [open, setOpen] = React.useState(false);
    const [myOpen, setMyOpen] = React.useState(false);
    const [pubOpen, setPubOpen] = React.useState(false);



    const handleClick = () => {
        setOpen(!open);
        if (!open) {
            setPubOpen(false);
            setMyOpen(false);
        }

    };
    const handlePubClick = () => {
        setPubOpen(!pubOpen);
        if (!pubOpen) {
            setOpen(false);
            setMyOpen(false);
        }
    };
    const handleMyClick = () => {
        setMyOpen(!myOpen);
        if (!myOpen) {
            setOpen(false);
            setPubOpen(false);
        }
    };



    return (
        <div>

            <Card className={classes.root}>
                <ListItem button onClick={handleClick}>
                    <ListItemIcon>
                        <AddBoxIcon/>
                    </ListItemIcon>
                    <ListItemText primary="Create Room"/>
                    {open ? <ExpandLess/> : <ExpandMore/>}
                </ListItem>
                <Collapse in={open} timeout="auto" unmountOnExit>
                    <CreateRoomForm/>
                </Collapse>

                <Divider/>

                <ListItem button onClick={handleMyClick}>
                    <ListItemIcon>
                        <AddBoxIcon/>
                    </ListItemIcon>
                    <ListItemText primary="My Room"/>
                    {myOpen ? <ExpandLess/> : <ExpandMore/>}
                </ListItem>
                <Collapse in={myOpen} timeout="auto" unmountOnExit>
                    <List>
                        <ListItem>
                            <FindMyRooms p={props}/>
                        </ListItem>
                    </List>
                </Collapse>

                <Divider/>

                <ListItem button onClick={handlePubClick}>
                    <ListItemIcon>
                        <AddBoxIcon/>
                    </ListItemIcon>
                    <ListItemText primary="Public Room"/>
                    {open ? <ExpandLess/> : <ExpandMore/>}
                </ListItem>
                <Collapse in={pubOpen} timeout="auto" unmountOnExit>
                    <FindPublicRooms/>
                </Collapse>
                <Divider/>
                <SignOut/>
            </Card>
        </div>


    )

}