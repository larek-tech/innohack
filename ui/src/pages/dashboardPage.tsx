import { LineChart, Line, BarChart, Bar, Tooltip, XAxis, YAxis, Legend, ResponsiveContainer } from 'recharts';

import { AppSidebar } from '@/components/app-sidebar';
import DashBoardService from '@/api/DashBoardService';
import { useEffect, useState } from 'react';
import { ChartReport, Record, Multiplier } from '@/api/models';
import { DatePicker } from '@/components/datepicker';

import { Button } from '@/components/ui/button';
import Markdown from 'react-markdown';


interface MultiplierCardProps {
    multiplier: Multiplier;
}
// @ts-ignore
const MultiplierCard: React.FC<MultiplierCardProps> = ({ multiplier }) => {
    return (
        <div className="bg-white shadow-md rounded-lg p-6 flex flex-col items-center justify-center">
            <h3 className="text-xl font-semibold mb-2 text-center">{multiplier.key}</h3>
            <p className="text-3xl font-bold text-blue-600">{multiplier.value}</p>
        </div>
    );
};

interface MyChartProps {
    title: string;
    legend: string[];
    data: any[];
    colors: string[];
    config?: any;
}



function transformChartReportToMyChartProps(chartReport: ChartReport): MyChartProps[] {
    const availableColors = [
        "#2563eb", // Blue
        "#60a5fa", // Light Blue
        "#DA3832", // Red
        "#34D399", // Green
        "#FBBF24", // Yellow
        "#A78BFA", // Purple
        "#F87171", // Light Red
        "#10B981", // Emerald
        "#F59E0B", // Amber
        "#8B5CF6", // Violet
    ];

    const graphs: MyChartProps[] = [];
    Object.keys(chartReport.info).forEach((key) => {
        console.log(chartReport.info[key], key);
        const leg = Object.keys(chartReport.info[key].legend)
        const data = new Map();
        for (let i = 0; i < chartReport.info[key].charts.length; i++) {

            chartReport.info[key].charts[i].records.forEach((record: Record) => {
                if (data.has(record.x)) {
                    data.set(record.x, [...data.get(record.x), record.y]);
                } else {
                    data.set(record.x, [record.y]);
                }
            });
        }
        const unpackedData = [...data.entries()].map(([x, yValues]) => {
            const obj: any = { x };
            // @ts-ignore
            yValues.forEach((y, index) => {
                obj[leg[index]] = y;
            });
            return obj;
        });
        const config = {}

        leg.forEach((key, index) => {
            // @ts-ignore
            config[key] = {
                color: availableColors[index % availableColors.length],
                label: key
            }
        });
        ;
        graphs.push({
            title: key,
            legend: leg,
            data: unpackedData,
            colors: leg.map((_, index) => availableColors[index % availableColors.length]),
        });

    });


    return graphs;
}

// @ts-ignore
export function BarChartComponent({ title, legend, data, colors, config }: MyChartProps) {
    return (
        <div className="border p-4 rounded-lg shadow-md flex flex-col" style={{ height: '300px' }}>
            <h2 className="text-center mb-4">{title}</h2>
            <ResponsiveContainer width="100%" height="100%">
                <BarChart data={data}>
                    <XAxis dataKey="x" />
                    <YAxis />
                    <Tooltip content={<CustomTooltip />} />
                    <Legend />
                    {legend.map((key, index) => (
                        <Bar key={index} dataKey={key} fill={colors[index]} radius={[4, 4, 0, 0]} />
                    ))}
                </BarChart>
            </ResponsiveContainer>
        </div>
    );
}// @ts-ignore
export function LineChartComponent({ title, legend, data, colors, config }: MyChartProps) {
    return (
        <div className="border p-4 rounded-lg shadow-md flex flex-col" style={{ height: '300px' }}>
            <h2 className="text-center mb-4">{title}</h2>
            <ResponsiveContainer width="100%" height="100%">
                <LineChart data={data}>
                    <XAxis dataKey="x" />
                    <YAxis />
                    <Tooltip content={<CustomTooltip />} />
                    <Legend />
                    {legend.map((key, index) => (
                        <Line
                            key={index}
                            type="monotone"
                            dataKey={key}
                            stroke={colors[index]}
                            strokeWidth={2}
                            dot={false}
                        />
                    ))}
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
}

const CustomTooltip = ({ active, payload, label }: any) => {
    if (active && payload && payload.length) {
        return (
            <div className="custom-tooltip bg-white p-2 border rounded shadow">
                <p className="label">{`Year: ${label}`}</p>
                {payload.map((entry: any, index: number) => (
                    <p key={`item-${index}`} style={{ color: entry.color }}>
                        {`${entry.name}: ${entry.value}`}
                    </p>
                ))}
            </div>
        );
    }
    return null;
};

export const DashBoardPage = () => {
    const [chartsData, setChartsData] = useState<ChartReport | null>(null);
    const [fromDate, setFromDate] = useState<Date>(new Date('2022-01-01'));
    const [toDate, setToDate] = useState<Date>(new Date('2023-01-01'));


    useEffect(() => {
        fetchReport();
    }, []);

    function fetchReport() {
        if (!fromDate || !toDate) {
            return;
        }
        DashBoardService.getReport({ startDate: fromDate, endDate: toDate }).then((data) => {
            setChartsData(data);
        });
    }



    return (
        <div className="flex w-full h-full">

            <AppSidebar />
            <div className="flex-1 p-4 m-y-4">
                <h1></h1>

                <div className="flex justify-between">
                    <DatePicker
                        date={fromDate}
                        // @ts-ignore
                        setDate={setFromDate}
                        title="From Date"
                    />
                    <h2> Финансовая аналитика </h2>
                    <DatePicker
                        date={toDate}
                        // @ts-ignore
                        setDate={setToDate}
                        title="To Date"
                    />
                </div>

                <div>
                    <Button onClick={fetchReport}>Построить отчет</Button>
                </div>
                <div>
                    <Markdown></Markdown>
                </div>
                <div>
                </div>
                {chartsData && fromDate === toDate && (
                    <><div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {transformChartReportToMyChartProps(chartsData).map((chartProps, index) => (
                            <BarChartComponent key={index} {...chartProps} />
                        ))}
                    </div><div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                            {transformChartReportToMyChartProps(chartsData).map((chartProps, index) => (
                                <LineChartComponent key={index} {...chartProps} />
                            ))}
                        </div></>
                )}
                <div className="m-2 p-8 border border-gray-300 rounded-lg shadow">
                    <Markdown>{chartsData?.summary}</Markdown>
                </div>
                {/* Charts Section */}
                {chartsData && (
                    <>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
                            {transformChartReportToMyChartProps(chartsData).map((chartProps, index) => (
                                <div key={index} className="max-w-full">
                                    {/* Render Bar and Line charts */}
                                    <BarChartComponent {...chartProps} />
                                    <LineChartComponent {...chartProps} />
                                </div>
                            ))}
                        </div>
                    </>
                )}
            </div>

        </div>
    )
}