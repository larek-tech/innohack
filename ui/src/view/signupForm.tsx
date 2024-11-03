import * as React from "react"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link, useNavigate } from "@tanstack/react-router"
import { useState } from "react"
import { useToast } from "@/hooks/use-toast"
import { useAuth } from "@/auth"
import { LoaderButton } from "@/components/ui/loader-button"


export function SignUp() {
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
        const passwordConfirm = formData.get("passwordConfirm") as string;
        if (password != passwordConfirm) {
            return
        }

        auth.signup({ email: email, password }, () => {
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
                    <CardTitle>Регистрация</CardTitle>
                </CardHeader>
                <CardContent>

                    <div className="grid w-full items-center gap-4">
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="email">Email</Label>
                            <Input id="email" name="email" placeholder="адрес электронной почты" />
                        </div>
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="email">Пароль</Label>
                            <Input id="password" name="password" type="password" placeholder="пароль" />
                        </div>
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="email">Подтверждение пароля</Label>
                            <Input id="password" name="passwordConfirm" type="password" placeholder="повторите свой пароль" />
                        </div>
                    </div>

                </CardContent>
                <CardFooter className="flex justify-between">
                    <Link
                        to="/login">Или авторизоваться</Link>
                    <LoaderButton isLoading={loading}>Зарегистрироваться</LoaderButton>
                </CardFooter>
            </form>
        </Card >
    )
}
