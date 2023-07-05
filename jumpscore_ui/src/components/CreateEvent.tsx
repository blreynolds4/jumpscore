import React from "react";
import { useRef, useState } from "react";
import Meet from "../models/Meet";

interface CreateEventProps {
  onNewMeet: (meet: Meet) => void;
}

const CreateEvent = ({ onNewMeet }: CreateEventProps) => {
  const [eventName, setEventName] = useState("");
  const [eventDate, setEventDate] = useState("");
  const dateInputRef = useRef(null);

  const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    console.log("using date", eventDate);
    onNewMeet({ name: eventName, date: new Date(eventDate) });
    setEventName("");
    setEventDate("");
  };

  const handleMeetNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEventName(event.target.value);
  };

  const handleMeetDateChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEventDate(event.target.value);
  };

  return (
    <div className="container-sm text-left">
      <h3>New Meet</h3>
      <form onSubmit={handleFormSubmit}>
        <div>
          <label>Meet Name:</label>
          <input
            value={eventName}
            type="text"
            onChange={handleMeetNameChange}
          />
        </div>
        <div>
          <label>Meet Date:</label>
          <input
            type="date"
            ref={dateInputRef}
            onChange={handleMeetDateChange}
          />
        </div>
        <button>Create Meet</button>
      </form>
    </div>
  );
};

export default CreateEvent;
