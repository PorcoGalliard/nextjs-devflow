import { popularTags, topQuestions } from "@/constants";
import React from "react";
import Image from "next/image";
import Link from "next/link";
import RenderTag from "./RenderTag";

const RightSideBar = () => {
  return (
    <section className="background-light900_dark200 custom-scrollbar light-border shadow-light-300 sticky right-0 top-0 flex h-screen flex-col justify-between overflow-y-auto border-l p-6 pt-36 dark:shadow-none max-xl:hidden lg:w-[350px] ">
      <div className="flex flex-col gap-1">
        <h3 className="text-dark200_light900 h3-bold max-lg:hidden">
          Top Questions
        </h3>
        <div className="mt-7 flex w-full flex-col gap-[30px]">
          {topQuestions.map((question) => (
            <Link
              key={question.label}
              href={`/questions/${question.label}`}
              className="text-dark200_light900 customer-pointer flex items-center justify-between gap-7"
            >
              <p className="text-dark200_light900 body-medium text-dark500_light700">
                {question.value}
              </p>
              <Image
                src="/assets/icons/chevron-right.svg"
                alt="chevron right"
                width={20}
                height={20}
                className="invert-colors"
              />
            </Link>
          ))}
        </div>
      </div>
      <div className="mt-16">
        <h3 className="h3-bold text-dark200_light900 max-lg:hidden">
          Popular Tags
        </h3>
        <div className="mt-7 flex flex-col gap-4">
          {popularTags.map((tag) => (
            <RenderTag
              key={tag._id}
              _id={tag._id}
              name={tag.name}
              totalQuestions={tag.totalQuestions}
              showCount
            />
          ))}
        </div>
      </div>
    </section>
  );
};

export default RightSideBar;
