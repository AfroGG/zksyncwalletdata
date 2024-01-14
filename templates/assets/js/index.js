const url = "https://mainnet.era.zksync.io";
let price = 0
const requestData = {
  jsonrpc: "2.0",
  id: 1,
  method: "zks_getTokenPrice",
  params: ["0x0000000000000000000000000000000000000000"],
};

fetch(url, {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify(requestData),
})
  .then((response) => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();
  })
  .then((data) => {
    price = data.result
    console.log(price)
    calculateAndDisplayTotal()
  })
  .catch((error) => {
    console.error("Error:", error);
  });


  function calculateAndDisplayTotal() {
    var totalGasSpentElement = document.getElementById("totalGasSpent");
    var totalsumValueElement = document.getElementById("sumValue");
    var totalnumElement = document.getElementById("num");
    var totaltxElement = document.getElementById("tx");

    let accountsdata = localStorage.getItem('accountsData');
    var accountsData = JSON.parse(accountsdata);
    let gas = 0;
    let valueSum = 0;
    let eth = 0;
    let usdc = 0;
    let txs = 0;
    const num = accountsData.length;

    for (let i = 0; i < num; i++) {
        gas = gas + accountsData[i].totalFee;
        eth = eth + accountsData[i].ethBalance;
        usdc = usdc + accountsData[i].usdcBalance;
        txs = txs + accountsData[i].nonce;
    }
    gas = (gas * price).toFixed(2);
    console.log(gas);
    console.log(usdc);
    console.log(eth);
    console.log(price);
    valueSum = ((eth * (+price)) + usdc).toFixed(2);;
    totalGasSpentElement.textContent = "$" + gas;
    totalsumValueElement.textContent = valueSum;
    totalnumElement.textContent = num;
    totaltxElement.textContent = txs;
}
document.addEventListener('DOMContentLoaded', function() {
let accountsdata = localStorage.getItem('accountsData');
var accountsData = JSON.parse(accountsdata);
const num = accountsData.length;
var bar1 = 0;
var bar2 = 0;
var bar3 = 0;
var bar4 = 0;
var bar5 = 0;
var bar6 = 0;
var bar7 = 0;
var bar8 = 0;
var bar9 = 0;
var bar10 = 0;

for (let i = 0; i < num; i++) {
			if (accountsData[i].nonce < 20) {
			  bar1++;
			} else if (accountsData[i].nonce < 30) {
			  bar2++;
			} else if (accountsData[i].nonce < 40) {
			  bar3++;
			} else if (accountsData[i].nonce < 50) {
			  bar4++;
			} else if (accountsData[i].nonce < 60) {
			  bar5++;
			} else if (accountsData[i].nonce < 70) {
			  bar6++;
			} else if (accountsData[i].nonce < 80) {
			  bar7++;
			} else if (accountsData[i].nonce < 90) {
			  bar8++;
			} else if (accountsData[i].nonce < 100) {
			  bar9++;
			} else if (accountsData[i].nonce < 110) {
			  bar10++;
			} else {
			  bar11++; // for nonce > 110
			}
		  }

var ctx = document.getElementById("myChart").getContext("2d");
var myChart = new Chart(ctx, {
    "type":"bar",
    "data":{
        "labels":[
            "0-20",
            "20-30",
            "30-40",
            "40-50",
            "50-60",
            "60-70",
            "70-80",
            "80-90",
            "90-100",
            "110-more"
        ],
        "datasets":[
            {
                "label":"count",
                "backgroundColor":"rgb(54,185,204)",
                "borderColor":"rgb(21,21,21)",
                "data":[
                    ""+bar1,
                    ""+bar2,
                    ""+bar3,
                    ""+bar4,
                    ""+bar5,
                    ""+bar6,
                    ""+bar7,
                    ""+bar8,
                    ""+bar9,
                    ""+bar10,
                    ""
                ]
            }
        ]
    },
    "options":{
        "maintainAspectRatio":false,
        "legend":{
            "display":false,
            "labels":{
                "fontStyle":"normal",
                "fontColor":"#000000"
            },
            "position":"top"
        },
        "title":{
            "fontStyle":"normal",
            "fontColor":"#000000",
            "display":true,
            "position":"top",
            "text":"ByNonce"
        },
        "scales":{
            "xAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ],
                        "drawOnChartArea":false
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ],
            "yAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ]
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ]
        }
    }
})


var bar11 = 0;
var bar12 = 0;
var bar13 = 0;
var bar14 = 0;
var bar15 = 0;
var bar16 = 0;
var bar17 = 0;

for (let i = 0; i < num; i++) {
			if (accountsData[i].activeWeek < 5) {
              bar11++;
			} else if (accountsData[i].activeWeek < 10) {
			  bar12++;
			} else if (accountsData[i].activeWeek < 15) {
			  bar13++;
			} else if (accountsData[i].activeWeek < 20) {
			  bar14++;
			} else if (accountsData[i].activeWeek < 25) {
			  bar15++;
			} else if (accountsData[i].activeWeek < 30) {
			  bar16++;
			} else {
			  bar17++; // for nonce > 110
			}
}


var ctx1 = document.getElementById("myChart1").getContext("2d");
var myChart1 = new Chart(ctx1, {
    "type":"bar",
    "data":{
        "labels":[
            "0-5",
            "5-10",
            "10-15",
            "15-20",
            "20-25",
            "25-30",
            "30-more"
        ],
        "datasets":[
            {
                "label":"count",
                "backgroundColor":"rgb(28,200,138)",
                "borderColor":"rgb(21,21,21)",
                "data":[
                    ""+bar11,
                    ""+bar12,
                    ""+bar13,
                    ""+bar14,
                    ""+bar15,
                    ""+bar16,
                    ""+bar17
                ]
            }
        ]
    },
    "options":{
        "maintainAspectRatio":false,
        "legend":{
            "display":false,
            "labels":{
                "fontStyle":"normal",
                "fontColor":"#000000"
            },
            "position":"top"
        },
        "title":{
            "fontStyle":"normal",
            "fontColor":"#000000",
            "display":true,
            "position":"top",
            "text":"ByActiveWeeks"
        },
        "scales":{
            "xAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ],
                        "drawOnChartArea":false
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ],
            "yAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ]
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ]
        }
    }
})


var bar21 = 0;
var bar22 = 0;
var bar23 = 0;
var bar24 = 0;
var bar25 = 0;
var bar26 = 0;
var bar27 = 0;
var bar28 = 0;
var bar29 = 0;
var bar30 = 0;
var bar31 = 0;
var bar32 = 0;
var bar33 = 0;
var bar34 = 0;
var bar35 = 0;

for (let i = 0; i < num; i++) {

			if (+accountsData[i].totalTxValue < 1000) {
              bar21++;
			} else if (+accountsData[i].totalTxValue < 2000) {
			  bar22++;
			} else if (+accountsData[i].totalTxValue < 3000) {
			  bar23++;
			} else if (+accountsData[i].totalTxValue < 4000) {
			  bar24++;
			} else if (+accountsData[i].totalTxValue < 5000) {
			  bar25++;
			} else if (+accountsData[i].totalTxValue < 6000) {
			  bar26++;
            } else if (+accountsData[i].totalTxValue < 7000) {
              bar27++;
			} else if (+accountsData[i].totalTxValue < 8000) {
			  bar28++;
			} else if (+accountsData[i].totalTxValue < 9000) {
			  bar29++;
			} else if (+accountsData[i].totalTxValue < 10000) {
			  bar30++;
            } else if (+accountsData[i].totalTxValue < 12000) {
			  bar31++;
			} else if (+accountsData[i].totalTxValue < 15000) {
			  bar32++;
			} else if (+accountsData[i].totalTxValue < 20000) {
			  bar33++;
			} else if (+accountsData[i].totalTxValue < 30000) {
			  bar34++;
			} else {
			  bar35++; // for nonce > 110
			}
}

var ctx2 = document.getElementById("myChart2").getContext("2d");
var myChart2 = new Chart(ctx2, {
    "type":"bar",
    "data":{
        "labels":[
            "less than 1k",
            "2k",
            "3k",
            "4k",
            "5k",
            "6k",
            "7k",
            "8k",
            "9k",
            "10k",
            "12k",
            "15k",
            "20k",
            "30k",
            "40k"
        ],
        "datasets":[
            {
                "label":"count",
                "backgroundColor":"rgb(78,115,223)",
                "borderColor":"rgb(21,21,21)",
                "data":[
                    ""+bar21,
                    ""+bar22,
                    ""+bar23,
                    ""+bar24,
                    ""+bar25,
                    ""+bar26,
                    ""+bar27,
                    ""+bar28,
                    ""+bar29,
                    ""+bar30,
                    ""+bar31,
                    ""+bar32,
                    ""+bar33,
                    ""+bar34,
                    ""+bar35
                ]
            }
        ]
    },
    "options":{
        "maintainAspectRatio":false,
        "legend":{
            "display":false,
            "labels":{
                "fontStyle":"normal",
                "fontColor":"#000000"
            },
            "position":"top"
        },
        "title":{
            "fontStyle":"normal",
            "fontColor":"#000000",
            "display":true,
            "position":"top",
            "text":"ByTxValue"
        },
        "scales":{
            "xAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ],
                        "drawOnChartArea":false
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ],
            "yAxes":[
                {
                    "gridLines":{
                        "color":"rgb(0,0,0)",
                        "zeroLineColor":"rgb(0,0,0)",
                        "drawBorder":true,
                        "drawTicks":false,
                        "borderDash":[
                            "2"
                        ],
                        "zeroLineBorderDash":[
                            "2"
                        ]
                    },
                    "ticks":{
                        "fontColor":"#141415",
                        "fontSize":"11",
                        "fontStyle":"normal",
                        "beginAtZero":false,
                        "padding":20
                    }
                }
            ]
        }
    }
})


var scoreCounts = {};
for (let i = 0; i < num; i++) {
    // console.log(Date.parse(accountsData[i].endDay - currentTimestamp))
    var account = accountsData[i];
    var score = account.score;

    // 检查 score 是否已经在 scoreCounts 中存在
    if (scoreCounts[score] === undefined) {
        // 如果不存在，则初始化为 1
        scoreCounts[score] = 1;
    } else {
        // 如果已存在，则递增计数
        scoreCounts[score]++;
    }
}
console.log(scoreCounts);

// 生成 data 部分的数据
var data = {
    labels: Object.keys(scoreCounts), // 获取所有分数
    datasets: [{
        label: "111", // 可以设置一个适当的标签
        backgroundColor: ["#1cc88a", "#4e73df", "#36b9cc", "#f6c23e", "#e74a3b", "#858796", "#6e707e",
        "#5a5c69", "#343a40", "#6a0dad", "#d7263d", "#3d84b8", "#f64f3c", "#5d2e46", "#3498db"], // 设置颜色
        borderColor: ["#ffffff", "#ffffff"],
        data: Object.values(scoreCounts), // 获取每个分数对应的对象数量
    }],
};

var ctx3 = document.getElementById("myChart3").getContext("2d");
var myChart3 = new Chart(ctx3, {
    type: "pie",
    data: data, // 使用生成的数据
    options: {
        maintainAspectRatio: false,
        legend: {
            display: false,
            labels: {
                fontStyle: "normal",
            },
        },
        title: {
            display: true, // 显示标题
            text: "Score Distribution", // 标题文本
            fontStyle: "normal",
        },
        plugins: {
            datalabels: { // 启用datalabels插件
                color: "white", // 标签颜色
                font: { size: 12 }, // 字体大小
                formatter: function(value, context) {
                    var label = context.chart.data.labels[context.dataIndex];
                    return label + ": " + value; // 显示标签和值
                }
            }
        },
    },
});








const currentTimestamp = Date.now();
var activeCount = 0;
var inActiveCount = 0;
for (let i = 0; i < num; i++) {
    const endDayTimestamp = Date.parse(accountsData[i].endDay);
    const timeDifference = currentTimestamp - endDayTimestamp;
    const oneWeekMilliseconds = 86400000 * 7;
    if (timeDifference > oneWeekMilliseconds) {
        inActiveCount++;
    } else {
        activeCount++;
    }
}
var ctx4 = document.getElementById("myChart4").getContext("2d");
var myChart4 = new Chart(ctx4, {
    "type":"pie",
    "data":{
        "labels":[
            "active",
            "inActive"
        ],
        "datasets":[
            {
                "label":"",
                "backgroundColor":[
                    "#1cc88a",
                    "#f70a0a"
                ],
                "borderColor":[
                    "#ffffff",
                    "#ffffff"
                ],
                "data":[
                    ""+activeCount,
                    ""+inActiveCount
                ]
            }
        ]
    },
    "options":{
        "maintainAspectRatio":false,
        "legend":{
            "display":false,
            "labels":{
                "fontStyle":"normal"
            }
        },
        "title":{
            "fontStyle":"normal"
        }
    }
})



});