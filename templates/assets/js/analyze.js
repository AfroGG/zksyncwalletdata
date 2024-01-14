document.addEventListener('DOMContentLoaded', function() {

    var loadingModal = document.getElementById('loadingModal');
    
    // 显示加载框
    function showLoadingModal() {
        loadingModal.style.display = 'block';
    }
    
    // 隐藏加载框
    function hideLoadingModal() {
        loadingModal.style.display = 'none';
    }
    
    var analyzeButton = document.getElementById('analyzeButton');
        
    analyzeButton.addEventListener('click', function(e) {
            e.preventDefault(); // 阻止按钮的默认行为
            console.log("提交表单")
            submitForm(); // 手动触发表单的提交
            
        });
        
        
    function submitForm() {
    
    var nonce = document.getElementById('nonce').value;
    var ethbalance = document.getElementById('ethbalance').value;
    var usdcbalance = document.getElementById('usdcbalance').value;
    var activeweeks = document.getElementById('activeweeks').value;
    var activedays = document.getElementById('activedays').value;
    var txvalue = document.getElementById('txvalue').value;
    var zklite_nonce = document.getElementById('zklite_nonce').value;
    var zklite_month = document.getElementById('zklite_month').value;
    var zklite_week = document.getElementById('zklite_week').value;
    var zklite_txvalue = document.getElementById('zklite_txvalue').value;
    var zklite_eth = document.getElementById('zklite_eth').value;
    var zklite_usdc = document.getElementById('zklite_usdc').value;
    
            // 使用对象字面量构建查询参数
    var queryParams = {
        nonce: nonce,
        ethbalance: ethbalance,
        usdcbalance: usdcbalance,
        activeweeks: activeweeks,
        activedays: activedays,
        txvalue: txvalue,
        zklite_nonce:zklite_nonce,
        zklite_month:zklite_month,
        zklite_week:zklite_week,
        zklite_txvalue:zklite_txvalue,
        zklite_eth:zklite_eth,
        zklite_usdc:zklite_usdc
     };
     showLoadingModal();
     localStorage.setItem('queryData', JSON.stringify(queryParams));
            // 构建完整的请求 URL
            var requestUrl = '/survey?' + new URLSearchParams(queryParams);
    
            fetch(requestUrl, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            })
            .then(function (response) {
                return response.json(); // 解析JSON响应
            })
            .then(function (data) {
                // 处理返回的JSON数据
                // 从localStorage中获取queryData的值
                hideLoadingModal();
                localStorage.setItem('data', JSON.stringify(data.data));
                localStorage.setItem('countTotal', JSON.stringify(data.countTotal));
                const countTotal = +localStorage.getItem('countTotal');
                console.log(data);
                window.location.reload();
                // renderDataToTable(nonce,ethbalance,usdcbalance,activeweeks,activedays,txvalue,zklite_nonce,zklite_month,zklite_week,zklite_txvalue,zklite_eth,zklite_usdc,countTotal);
            })
            .catch(function (error) {
                console.error('Error:', error);
            });
        }
        const countTotal = +localStorage.getItem('countTotal');
        renderDataToTable(nonce,ethbalance,usdcbalance,activeweeks,activedays,txvalue,zklite_nonce,zklite_month,zklite_week,zklite_txvalue,zklite_eth,zklite_usdc,countTotal);
        // window.location.reload();
    
    

    function renderDataToTable(nonce,ethbalance,usdcbalance,activeweeks,activedays,txvalue,zklite_nonce,zklite_month,zklite_week,zklite_txvalue,zklite_eth,zklite_usdc,countTotal){
        if (myChart5) {
            myChart5.destroy();
        }
        if (myChart6) {
            myChart6.destroy();
        }
    let count = +localStorage.getItem('data');
    var ctx5 = document.getElementById("myChart5").getContext("2d");
    var myChart5 = new Chart(ctx5, {
        type: "pie",
        data: {
            labels: ["count", "else"],
            datasets: [
                {
                    label: "",
                    backgroundColor: ["#1cc88a", "#f70a0a"],
                    borderColor: ["#ffffff", "#ffffff"],
                    data: [count, countTotal - count],
                },
            ],
        },
        options: {
            maintainAspectRatio: false,
            legend: {
                display: false,
                labels: {
                    fontStyle: "normal",
                },
            },
            title: {
                fontStyle: "normal",
            },
            plugins: {
                datalabels: {
                    display: true,
                    formatter: function (value, context) {
                        var dataset = context.chart.data.datasets[0];
                        var total = dataset.data.reduce(function (acc, current) {
                            return acc + current;
                        }, 0);
                        var percentage = ((value / total) * 100).toFixed(2) + "%";
                        return value + " (" + percentage + ")";
                    },
                },
            },
        },
    });
    
    const Data = localStorage.getItem('queryData');
    let queryData = JSON.parse(Data);
    const address = localStorage.getItem('accountsData');
    let accountsData = JSON.parse(address);
       console.log(accountsData)
    var myWalletCount = 0
      for (let i=0;i < accountsData.length;i++){
        if (accountsData[i].activeDay>=queryData.activedays&&accountsData[i].activeWeek>=queryData.activeweeks&&accountsData[i].zklite_nonce>=queryData.zklite_nonce&&accountsData[i].zklite_month>=queryData.zklite_month&&accountsData[i].totalTxValue>=queryData.txvalue&&accountsData[i].ethBalance>=queryData.ethbalance&&accountsData[i].usdcBalance>=queryData.usdcbalance&&accountsData[i].zklite_week>=queryData.zklite_week&&accountsData[i].zklite_txvalue>=queryData.zklite_txvalue&&accountsData[i].zklite_eth>=queryData.zklite_eth&&accountsData[i].zklite_usdc>=queryData.zklite_usdc){
            myWalletCount++
        }}
    console.log(myWalletCount)
    var ctx6 = document.getElementById("myChart6").getContext("2d");
    var myChart6 = new Chart(ctx6, {
        "type":"pie",
        "data":{
            "labels":[
                "count",
                "else"
            ],
            "datasets":[
                {
                    "label":"",
                    "backgroundColor":[
                        "#1cc88a",
                        "#ADD8E6"
                    ],
                    "borderColor":[
                        "#ffffff",
                        "#ffffff"
                    ],
                    "data":[
                        ""+myWalletCount,
                        ""+(accountsData.length-myWalletCount)
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
    var queryDataDisplay = document.getElementById('queryDataDisplay');
    var queryDataHTML = '<h6>Query Data:</h6>';
        queryDataHTML += '<ul>';
        for (var key in queryData) {
            queryDataHTML += '<li>' + key + ': ' + queryData[key] + '</li>';
        }
        queryDataHTML += '</ul>';
        queryDataDisplay.innerHTML = queryDataHTML;

    const DataCount = localStorage.getItem('data');
    const TotalCount = localStorage.getItem('countTotal');
    var DataDisplay = document.getElementById('DataDisplay');
    var DataHTML = '<h6>Data</h6>';
    DataHTML += '<ul>';
    DataHTML += '<li>Total Wallets: ' + TotalCount + '</li>';
    DataHTML += '<li>Eligible Wallets: ' + DataCount + '</li>';
    DataHTML += '<li>Other Wallets: ' + (TotalCount - DataCount) + '</li>';
    DataHTML += '</ul>';
    DataDisplay.innerHTML = DataHTML;
    
    }})

    