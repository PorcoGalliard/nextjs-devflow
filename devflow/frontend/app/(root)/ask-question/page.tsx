"use client";

import Question from "@/components/forms/Question";
import User from "@/database/user.model";
import axios from "axios";
import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";

const Page = () => {
  const router = useRouter();
  const userID = "12345678";
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    if (!userID) router.push("/sign-in");

    const fetchUser = async () => {
      const response = await axios.get(
        `http://localhost:5000/api/v1/user/${userID}`
      );
      const mongoUser = response.data;
      console.log(mongoUser);
      setUser(mongoUser);
    };

    fetchUser();
  }, []);

  if (!user) return <div>Loading...</div>;

  return (
    <div>
      <h1 className="h1-bold text-dark100_light900">Ask a Question</h1>
      <div className="mt-9">
        <Question mongoUserId={JSON.stringify(user.clerkID)} />
      </div>
    </div>
  );
};

export default Page;
