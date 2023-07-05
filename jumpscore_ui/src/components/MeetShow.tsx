import React from "react";
import Meet from "../models/Meet";

interface MeetShowProps {
  meet: Meet;
  onDelete: (name: string) => void;
}

function MeetShow({ meet, onDelete }: MeetShowProps) {
  const handleClick = () => {
    onDelete(meet.name);
  };
  return (
    <li className="list-group-item d-flex justify-content-between align-items-start">
      <div className="ms-2 me-auto">
        <div className="fw-bold">{meet.name}</div>
        {meet.date.toLocaleDateString()}
      </div>
      <button type="button" onClick={handleClick} className="btn-close" />
    </li>
  );
}

export default MeetShow;
