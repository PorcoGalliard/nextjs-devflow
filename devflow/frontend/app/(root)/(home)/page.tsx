"use client";

import QuestionCard from "@/components/card/QuestionCard";
import Question from "@/database/question.model";
import HomeFilters from "@/components/home/HomeFilters";
import Filter from "@/components/shared/Filter";
import NoResult from "@/components/shared/NoResult";
import LocalSearch from "@/components/shared/search/LocalSearch";
import { Button } from "@/components/ui/button";
import { HomePageFilters } from "@/constants/filters";
import Link from "next/link";
import { useEffect, useState } from "react";
import axios from "axios";

export default function Home() {
  const [questions, setQuestions] = useState<Question[]>([]);

  useEffect(() => {
    const fetchQuestions = async () => {
      const response = await axios.get("http://localhost:5000/api/v1/question");
      const mongoQuestions = response.data;
      console.log(mongoQuestions);
      setQuestions(mongoQuestions);
    };

    fetchQuestions();
  }, []);

  if (!questions)
    return (
      <div>
        {" "}
        <NoResult
          title="There is no question to show"
          description="        Be the first to break the silence! 🚀 Ask a Question and kickstart the
discussion. our query could be the next big thing others learn from. Get
involved! 💡"
          link="/ask-question"
          linkTitle="Ask a Question"
        />
      </div>
    );

  return (
    <>
      <div className="flex w-full flex-col-reverse justify-between gap-4 sm:flex-row sm:items-center">
        <h1 className="h1-bold text-dark100_light900">All Questions</h1>
        <Link href="/ask-question" className="flex justify-end max-sm:w-full">
          <Button className="!text-light-900 primary-gradient shadow-light-700 min-h-[46px] px-4 py-3">
            Ask a Question
          </Button>
        </Link>
      </div>
      <div className="mt-11 flex justify-between gap-5 max-sm:flex-col sm:items-center ">
        <LocalSearch
          placeholder="Search questions..."
          route="/"
          iconPosition="left"
          imgSrc="/assets/icons/search.svg"
          otherClasses="flex-1"
        />
        <Filter
          filters={HomePageFilters}
          otherClasses="min-h-[56px] sm:min-w-[170px]"
          containerClasses="hidden max-md:flex"
        />
      </div>
      <HomeFilters />
      <div className="mt-10 flex w-full flex-col gap-6">
        {questions.map((question) => (
          <QuestionCard
            key={question._id}
            _id={question._id}
            user={question.user}
            title={question.title}
            tags={question.tagDetails}
            upvotes={question.upvotes}
            views={question.views}
            answers={question.answers}
            createdAt={question.createdAt}
          />
        ))}
      </div>
    </>
  );
}
