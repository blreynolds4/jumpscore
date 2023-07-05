import { useState } from "react";
import CreateEvent from "./components/CreateEvent";
import Meet from "./models/Meet";
import getEvents from "./api";
import MeetList from "./components/MeetList";

function App() {
  const [meetList, setMeetList] = useState([]);

  const createMeet = (meet: Meet) => {
    console.log("handling new meet", meet.name, meet.date);
    const updatedMeets = [...meetList, meet];
    setMeetList(updatedMeets);
    console.log("MEETS", updatedMeets);
  };

  const deletMeetByName = (name: string) => {
    console.log("DELETING ", name);
    const updatedMeets = meetList.filter((meet) => {
      return meet.name != name;
    });
    setMeetList(updatedMeets);
  };

  return (
    <div className="container-sm text-center">
      <h1>Ski Jump Meet Scoring</h1>
      Meet Count: {meetList.length}
      <MeetList meets={meetList} onDelete={deletMeetByName} />
      <CreateEvent onNewMeet={createMeet} />
    </div>
  );
}

export default App;
