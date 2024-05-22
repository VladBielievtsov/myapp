"use client";

import React from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { Soup } from "lucide-react";
import Button from "@/components/common/Button";
import Input from "@/components/common/Input";

type Inputs = {
  username: string;
  password: string;
};

export default function LoginForm() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = (data) => console.log(data);

  return (
    <div className="max-w-[400px] w-full">
      <div className="text-zinc-800 dark:text-zinc-300 flex flex-col items-center">
        <Soup size={32} strokeWidth={2} />
        <span className="text-2xl py-[30px] font-bold">myapp</span>
      </div>
      <form
        className="flex flex-col max-w-[300px] w-full mx-auto gap-9"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div>
          <div className="mb-2.5">
            <label className="text-sm font-bold text-zinc-800 dark:text-zinc-300">
              Username
            </label>
          </div>
          <div>
            <Input
              type="text"
              defaultValue="admin"
              {...register("username", { required: true })}
            />
            {errors.username && (
              <div className="mt-2.5">
                <span className="text-sm font-bold text-red-500">Required</span>
              </div>
            )}
          </div>
        </div>
        <div>
          <div className="mb-2.5">
            <label className="text-sm font-bold text-zinc-800 dark:text-zinc-300">
              Password
            </label>
          </div>
          <div>
            <Input
              type="password"
              {...register("password", { required: true })}
            />
            {errors.password && (
              <div className="mt-2.5">
                <span className="text-sm font-bold text-red-500">Required</span>
              </div>
            )}
          </div>
        </div>
        <div>
          <Button type="submit" className="w-full min-h-[42px]">
            Login
          </Button>
        </div>
      </form>
    </div>
  );
}
