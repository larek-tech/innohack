import * as React from "react"
import { useState } from 'react';
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link, useNavigate } from "@tanstack/react-router"
import { API_URL } from '@/config';
import { useAuth } from "@/auth"
import { useToast } from "@/hooks/use-toast"
import { LoaderButton } from "@/components/ui/loader-button";

export function LoginForm() {
    const navigate = useNavigate();
    const auth = useAuth();
    const { toast } = useToast();

    const [loading, setLoading] = useState<boolean>(false);

    const from = location.state?.from?.pathname || '/';

    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget);
        const email = formData.get('email') as string;
        const password = formData.get('password') as string;

        setLoading(true);

        auth.login({ email: email, password }, () => {
            navigate({ from, to: "/" })
        })
            .catch(() => {
                toast({
                    title: 'Ошибка',
                    description: 'Ошибка аутентификации. Попробуйте еще раз.',
                    variant: 'destructive',
                });
            })
            .finally(() => {
                setLoading(false);
            });
    }

    return (
        <div className="flex flex-col md:flex-row min-h-screen">
            <div className="hidden md:flex md:w-1/2 bg-gray-100 dark:bg-gray-800 items-center justify-center">
                <h1 className="text-4xl font-bold">Welcome to ФинансовыйПоиск</h1>
            </div>
            <div className="flex w-full md:w-1/2 items-center justify-center p-4">
                <Card className="w-full max-w-md">
                    <form onSubmit={handleSubmit}>
                        <CardHeader>
                            <CardTitle>Авторизация</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <div className="grid w-full items-center gap-4">
                                <div className="flex flex-col space-y-1.5">
                                    <Label htmlFor="email">Email</Label>
                                    <Input id="email" name="email" type="email" placeholder="адрес электронной почты" autoComplete="email" />
                                </div>
                                <div className="flex flex-col space-y-1.5">
                                    <Label htmlFor="password">Password</Label>
                                    <Input id="password" name="password" placeholder="пароль" type="password" autoComplete="current-password" />
                                </div>
                            </div>
                        </CardContent>
                        <CardFooter className="flex justify-between">
                            <Link to="/signup">Или зарегистрироваться</Link>
                            <LoaderButton isLoading={loading} type="submit">Авторизоваться</LoaderButton>
                        </CardFooter>
                    </form>
                </Card>
            </div>
        </div>
    )
}