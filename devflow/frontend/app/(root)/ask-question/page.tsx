import Question from "@/components/forms/Question";
import { auth } from "@clerk/nextjs";
import axios from "axios";
import { redirect } from "next/navigation";
import React from "react";

const page = () => {
  const { userId } = auth();

  if (!userId) redirect("/sign-in");

  axios.get(`http://localhost:5000/api/v1/user/${userId}`).then((response) => {
    const mongoUser = response.data;
    console.log(mongoUser);
  });

  return (
    <div>
      <h1 className="h1-bold text-dark100_light900">Ask a Question</h1>
      <div className="mt-9">
        <Question />
      </div>
    </div>
  );
};

export default page;
