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

        <Card className="w-1/2">
            <form onSubmit={handleSubmit}>
                <CardHeader>
                    <CardTitle>Авторизация</CardTitle>
                    {/* <CardDescription>Deploy your new project in one-click.</CardDescription> */}
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
                    <Link
                        to="/signup">Или зарегистрироваться</Link>
                    <LoaderButton isLoading={loading} type="submit">Авторизоваться</LoaderButton>
                </CardFooter>
            </form >
        </Card >

    )
}
