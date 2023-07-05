import React from "react";
import MeetShow from "./MeetShow";

interface MeetListProps {
  meets: [];
  onDelete: (name: string) => void;
}

const MeetList = ({ meets, onDelete }: MeetListProps) => {
  const renderedMeets = meets.map((meet) => {
    return <MeetShow key={meet.name} meet={meet} onDelete={onDelete} />;
  });

  return <ul className="list-group">{renderedMeets}</ul>;
};

export default MeetList;
