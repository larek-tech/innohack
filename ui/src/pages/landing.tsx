import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Link } from "@tanstack/react-router"
import { Search, BarChart, PieChart, TrendingUp, Building } from "lucide-react"

export const Landing = () => {
    return (
        <div className="flex flex-col min-h-screen" >
            <header className="px-4 lg:px-6 h-14 flex items-center">
                <Link className="flex items-center justify-center" href="#">
                    <Building className="h-6 w-6" />
                    <span className="ml-2 text-lg font-bold">ФинансовыйПоиск</span>
                </Link>
                <nav className="ml-auto flex gap-4 sm:gap-6">
                    <Link className="text-sm font-medium hover:underline underline-offset-4" href="#">
                        Функции
                    </Link>
                    <Link className="text-sm font-medium hover:underline underline-offset-4" href="#">
                        О МТС
                    </Link>
                    <Link className="text-sm font-medium hover:underline underline-offset-4" href="#">
                        Контакты
                    </Link>
                </nav>
            </header>
            <main className="flex-1">
                <section className="w-full py-12 md:py-24 lg:py-32 xl:py-48">
                    <div className="container px-4 md:px-6">
                        <div className="flex flex-col items-center space-y-4 text-center">
                            <div className="space-y-2">
                                <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none">
                                    Исследуйте Финансовую Аналитику МТС
                                </h1>
                                <p className="mx-auto max-w-[700px] text-gray-500 md:text-xl dark:text-gray-400">
                                    Укрепите свой финансовый анализ с помощью нашей продвинутой системы поиска. Найдите полные данные по компании МТС за считанные секунды.
                                </p>
                            </div>
                            <div className="w-full max-w-sm space-y-2">
                                <form className="flex space-x-2">
                                    <Input className="max-w-lg flex-1" placeholder="Поиск финансовых данных МТС..." type="text" />
                                    <Button type="submit">
                                        <Search className="mr-2 h-4 w-4" />
                                        Поиск
                                    </Button>
                                </form>
                            </div>
                        </div>
                    </div>
                </section>
                <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-100 dark:bg-gray-800">
                    <div className="container px-4 md:px-6">
                        <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl text-center mb-12">Ключевые Функции</h2>
                        <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3">
                            <div className="flex flex-col items-center space-y-2 border-gray-800 p-4 rounded-lg">
                                <Search className="h-8 w-8 mb-2" />
                                <h3 className="text-xl font-bold">Продвинутый Поиск</h3>
                                <p className="text-sm text-gray-500 dark:text-gray-400 text-center">
                                    Мощные алгоритмы поиска для быстрого нахождения точной финансовой информации.
                                </p>
                            </div>
                            <div className="flex flex-col items-center space-y-2 border-gray-800 p-4 rounded-lg">
                                <BarChart className="h-8 w-8 mb-2" />
                                <h3 className="text-xl font-bold">Визуализация Данных</h3>
                                <p className="text-sm text-gray-500 dark:text-gray-400 text-center">
                                    Интерактивные диаграммы и графики для лучшего понимания данных.
                                </p>
                            </div>
                            <div className="flex flex-col items-center space-y-2 border-gray-800 p-4 rounded-lg">
                                <TrendingUp className="h-8 w-8 mb-2" />
                                <h3 className="text-xl font-bold">Обновления в Реальном Времени</h3>
                                <p className="text-sm text-gray-500 dark:text-gray-400 text-center">
                                    Будьте в курсе последних финансовых данных и рыночных трендов.
                                </p>
                            </div>
                        </div>
                    </div>
                </section>
                <section className="w-full py-12 md:py-24 lg:py-32">
                    <div className="container px-4 md:px-6">
                        <div className="grid gap-10 lg:grid-cols-2 items-center">
                            <div className="space-y-4">
                                <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl">О Компании МТС</h2>
                                <p className="max-w-[600px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
                                    МТС (Мобильные Телесистемы) — ведущая телекоммуникационная группа в России и СНГ. Наша платформа предоставляет полные финансовые данные и анализ, чтобы помочь вам принимать обоснованные решения о МТС и телекоммуникационном рынке.
                                </p>
                            </div>
                            <div className="flex justify-center">
                                <PieChart className="h-full w-full max-h-72" />
                            </div>
                        </div>
                    </div>
                </section>
                <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-100 dark:bg-gray-800">
                    <div className="container px-4 md:px-6">
                        <div className="flex flex-col items-center justify-center space-y-4 text-center">
                            <div className="space-y-2">
                                <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl">
                                    Начните Свой Финансовый Анализ Сегодня
                                </h2>
                                <p className="mx-auto max-w-[600px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
                                    Присоединяйтесь к тысячам аналитиков, которые доверяют нашей платформе для точной и своевременной финансовой аналитики МТС и телекоммуникационной отрасли.
                                </p>
                            </div>
                            <div className="w-full max-w-sm space-y-2">
                                <Link to="/signup">
                                    <Button className="w-full" size="lg">
                                        Начать

                                    </Button>
                                </Link>
                            </div>
                        </div>
                    </div>
                </section>
            </main>
            <footer className="flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t">
                <p className="text-xs text-gray-500 dark:text-gray-400">© 2024 ФинансовыйПоиск. Все права защищены.</p>
                <nav className="sm:ml-auto flex gap-4 sm:gap-6">
                    <Link className="text-xs hover:underline underline-offset-4" href="#">
                        Условия использования
                    </Link>
                    <Link className="text-xs hover:underline underline-offset-4" href="#">
                        Конфиденциальность
                    </Link>
                </nav>
            </footer>
        </div >
    )
}
