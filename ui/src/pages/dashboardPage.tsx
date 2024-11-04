import { BarChart, Bar, Tooltip, XAxis, YAxis, Legend } from 'recharts';
import { ChartConfig, ChartContainer } from "@/components/ui/chart"
import { AppSidebar } from '@/components/app-sidebar';
import DashBoardService from '@/api/DashBoardService';
import { useEffect, useState } from 'react';
import { Chart, ChartReport, Record, Info, Multiplier } from '@/api/models';
import { DatePicker } from '@/components/datepicker';
// import { Button as DayPickerButton, DayPickerProvider } from 'react-day-picker';
import { Button } from '@/components/ui/button';
// import { Sheet, SheetTrigger, SheetContent, SheetHeader, SheetTitle } from '@/components/ui/sheet';
// import { ScrollArea } from '@/components/ui/scroll-area';
// import { Input } from '@/components/ui/input';
// import { MessageCircle, Send } from 'lucide-react';
// import ChatInterface from './chatPage';



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
            // get all X values and put it as key to data map
            // then add all Y values to the corresponding X value if it exists otherwise insert 0
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
            yValues.forEach((y, index) => {
                obj[leg[index]] = y;
            });
            return obj;
        });
        const config = {}
        // for every element in leg create a new key in config and place there the color and label as value
        leg.forEach((key, index) => {
            config[key] = {
                color: availableColors[index % availableColors.length],
                label: key
            }
        });
        // console.log(leg, config);
        // console.log(unpackedData);
        graphs.push({
            title: key,
            legend: leg,
            data: unpackedData,
            colors: leg.map((_, index) => availableColors[index % availableColors.length]),
        });

    });


    return graphs;
}


export function BarChartComponent({ title, legend, data, colors, config }: MyChartProps) {
    // Deduplicate x axis values


    config = {
        "Коэффициент автономии": {
            "color": "#2563eb",
            "label": "Коэффициент автономии"
        },
        "Коэффициент капитализации": {
            "color": "#60a5fa",
            "label": "Коэффициент капитализации"
        },
        "Коэффициент обеспеченности материальных запасов": {
            "color": "#DA3832",
            "label": "Коэффициент обеспеченности материальных запасов"
        },
        "Коэффициент покрытия инвестиций": {
            "color": "#34D399",
            "label": "Коэффициент покрытия инвестиций"
        },
        "Коэффициент финансового левериджа": {
            "color": "#FBBF24",
            "label": "Коэффициент финансового левериджа"
        },
        "Коэффициент финансовой зависимости": {
            "color": "#A78BFA",
            "label": "Коэффициент финансовой зависимости"
        }
    }

    data = [
        [
            "2020",
            [
                0.06700287069506347,
                0.9133202009221758,
                0.7729929165491712,
                -706.4488597346816,
                0.7059900458541079,
                4.25276298517291
            ]
        ],
        [
            "2021",
            [
                0.06382820869815833,
                0.9049114103400254,
                0.6712499252160579,
                -988.4630302224223,
                0.6074217165178996,
                3.0923952600483853
            ]
        ],
        [
            "2022",
            [
                0.030574102133696067,
                0.9484523395286723,
                0.5931229827724628,
                -412.2476254578938,
                0.5625488806387667,
                1.946623032244732
            ]
        ],
        [
            "2023",
            [
                0.04328000532871856,
                0.9219393778432028,
                0.5544409477262912,
                -227.92772877401373,
                0.5111609423975726,
                1.1926627774694993
            ]
        ]
    ]
    console.log(data);
    data = data.map(([x, yValues]) => {
        const obj: any = { x };
        yValues.forEach((y, index) => {
            obj[legend[index]] = y;
        });
        return obj;
    });

    return (
        <ChartContainer config={config} className="min-h-[200px] w-full">
            <BarChart accessibilityLayer data={data}>
                {legend.map((key, index) => (
                    console.log(key, index),
                    <Bar key={index} dataKey={key} fill={colors[index]} radius={4} />
                ))}
                <Tooltip content={<CustomTooltip />} />
                <XAxis dataKey="x" />
                <YAxis />
                <Legend />
            </BarChart>
        </ChartContainer>
    )
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

export const DashBoardPage = () => {
    const [chartsData, setChartsData] = useState<ChartReport | null>(null);
    const [fromDate, setFromDate] = useState<Date>(new Date('2020-01-01'));
    const [toDate, setToDate] = useState<Date>(new Date('2023-12-31'));
    const [isChatOpen, setIsChatOpen] = useState(false);


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
            <AppSidebar className="w-1/4" />
            <div className="flex-1 p-4 m-y-4">
                <h1></h1>

                <div className="flex justify-between">
                    <DatePicker
                        date={fromDate}
                        setDate={setFromDate}
                        title="From Date"
                    />
                    <h2> Analysis Interval </h2>
                    <DatePicker
                        date={toDate}
                        setDate={setToDate}
                        title="To Date"
                    />
                </div>

                <div>
                    <Button onClick={fetchReport}>Get Report</Button>
                </div>
                {chartsData && (
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {transformChartReportToMyChartProps(chartsData).map((chartProps, index) => (
                            <BarChartComponent key={index} {...chartProps} />
                        ))}
                    </div>
                )}
            </div>

        </div>
    )
}