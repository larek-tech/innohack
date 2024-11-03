import { BarChart, Bar } from 'recharts';
import { ChartConfig, ChartContainer } from "@/components/ui/chart"
import { AppSidebar } from '@/components/app-sidebar';
import DashBoardService from '@/api/DashBoardService';
import { useEffect, useState } from 'react';
import { Chart, ChartReport, Record } from '@/api/models';
import { XAxis, YAxis } from 'recharts';

import { Legend } from 'recharts';
import { Tooltip } from '@/components/ui/tooltip';

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
]

function generateChartConfig(keys: string[], colors: string[]): ChartConfig {
    const config: ChartConfig = {};
    keys.forEach((key, index) => {
        config[key] = {
            label: key.charAt(0).toUpperCase() + key.slice(1),
            color: colors[index % colors.length],
        };
    });
    return config;
}

interface MyChartProps {
    data: any[];
    config: ChartConfig;
}

function convertData(data: ChartReport): Array<{ config: ChartConfig, [key: string]: any[] }> {
    const result: Array<{ config: ChartConfig, [key: string]: any[] }> = [];
    data.charts.forEach(chart => {
        const chartData = {
            config: generateChartConfig(Object.keys(chart.records), availableColors),
            [chart.description]: chart.records,
        };
        result.push(chartData);
    });
    return result;
}

const CustomTooltip = ({ active, payload, label }: any) => {
    if (active && payload && payload.length) {
        return (
            <div className="custom-tooltip">
                <p className="label">{`Key: ${label}`}</p>
                {payload.map((entry: any, index: number) => (
                    <p key={`item-${index}`} style={{ color: entry.color }}>{`${entry.name}: ${entry.value}`}</p>
                ))}
            </div>
        );
    }
    return null;
};



export function MyChart({ data, config }: MyChartProps) {
    return (
        <div className="border p-4 rounded-lg shadow-md">
            <h2>Data</h2>
            <ChartContainer config={config} className="min-h-[200px] w-full">
                <BarChart data={data}>
                    <XAxis dataKey="x" label={{ value: 'X Axis', position: 'insideBottomRight', offset: -5 }} />
                    <YAxis label={{ value: 'Y Axis', angle: -90, position: 'insideLeft' }} />
                    <Tooltip content={<CustomTooltip />} />
                    <Legend />
                    {Object.keys(config).map((key) => (
                        <Bar key={key} dataKey="y" name={config[key].label} fill={config[key].color} radius={4} />
                    ))}
                </BarChart>
            </ChartContainer>
        </div>
    )
}

export const DashBoardPage = () => {
    const [chartsData, setChartsData] = useState<Array<{ config: ChartConfig, [key: string]: any[] }>>([]);

    useEffect(() => {
        DashBoardService.getReport({ startDate: new Date('2020-01-01'), endDate: new Date('2023-12-31') }).then((data) => {
            const result = convertData(data);
            setChartsData(result);
        });
    }, []);

    return (
        <div className="flex">
            <AppSidebar className="w-1/4" />
            <div className="flex-1 p-4">
                <h1>Dash board</h1>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {chartsData.map((chart, index) => {
                        const [key, data] = Object.entries(chart).find(([k]) => k !== "config")!;
                        return (
                            <div key={index} className='w-96 h-96'>
                                <MyChart data={data} config={chart.config} />
                            </div>
                        );
                    })}
                </div>
            </div>
        </div>
    )
}