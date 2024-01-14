document.addEventListener("DOMContentLoaded", function () {
    var Updatingmodal = document.getElementById('Updatingmodal');
    
    // 显示加载框
    function showLoadingModal() {
        Updatingmodal.style.display = 'block';
    }
    
    // 隐藏加载框
    function hideLoadingModal() {
        Updatingmodal.style.display = 'none';
    }

    document.querySelector('.scroll-to-top').addEventListener('click', function(e) {
        e.preventDefault(); // 阻止默认的超链接行为
        const tableContainer = document.querySelector('.table-container');
        tableContainer.scrollTop = 0; // 将滚动条置顶
      });
    function isAddressExists(address) {
        const rows = document.querySelectorAll('tbody tr'); // 获取所有表格行
        for (const row of rows) {
            const existingAddress = row.querySelector('td:nth-child(2) div').textContent; // 假设地址在第二列
            if (existingAddress === address) {
                return true; // 地址已存在，返回true
            }
        }
        return false; // 地址不存在，返回false
    }
    function renderDataToTable(data) {
        const tbody = document.querySelector('tbody');
    
        // 检查 data.Accounts 是否存在
        if (data.Accounts && data.Accounts.length > 0) {
            // 存储 Accounts 数据到 localStorage
            localStorage.setItem('accountsData', JSON.stringify(data.Accounts));
    
            data.Accounts.forEach(function (account, index) {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td><input type="checkbox"></td>
                    <td class="ant-table-cell" style="text-align: center;"><div style="background-color: rgb(187, 238, 250); border-radius: 5px; width: 160px;font-size: 12px;overflow: hidden; text-overflow: ellipsis;word-wrap: break-word;"><a href="https://debank.com/profile/${account.address}" target="_blank">${account.address}</a></div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.ethBalance}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.usdcBalance}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.nonce}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.activeWeek}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.totalTxValue.toFixed(2)}</div></td>
                    <td class="ant-table-cell" style="text-align: center; background-color:pink; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.score}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.rank}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.totalFee}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zkMonth}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.startDay.split("T")[0]}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.endDay.split("T")[0]}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_nonce}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_txvalue.toFixed(2)}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_week}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_month}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_startDay.split("T")[0]}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_eth.toFixed(2)}</div></td>
                    <td class="ant-table-cell" style="text-align: center; vertical-align: middle;"><div style="color: rgb(0, 0, 0);">${account.zklite_usdc.toFixed(2)}</div></td>
                `;
                tbody.appendChild(row);
        
                // 添加行号
                const cell = document.createElement("td");
                cell.textContent = index + 1;
                row.insertBefore(cell, row.children[1]);
            });
    
}
}
const tbody = document.querySelector('tbody');
const deleteButton = document.querySelector('.delete-button');
deleteButton.addEventListener("click", function () {
    const checkedCheckboxes = document.querySelectorAll('input[type="checkbox"]:checked');
    if (checkedCheckboxes.length > 0) {
        const storedData = JSON.parse(localStorage.getItem('accountsData')) || [];
        
        // 倒序遍历选中的复选框，以避免索引问题
        for (let i = checkedCheckboxes.length - 1; i >= 0; i--) {
            const checkbox = checkedCheckboxes[i];
            // 检查是否是全选框所在的行
            if (!checkbox.closest('th:first-child')) {
                const rowNumber = parseInt(checkbox.closest('tr').querySelector('td:first-child + td').textContent);
                if (!isNaN(rowNumber)) {
                    // 删除表格中的行
                    const rowToDelete = tbody.querySelector(`tr:nth-child(${rowNumber})`);
                    if (rowToDelete) {
                        rowToDelete.remove();
                    }

                    // 删除本地存储中对应行号-1的数据
                    storedData.splice(rowNumber - 1, 1);
                }
            }
        }
        // 更新本地存储中的数据
        localStorage.setItem('accountsData', JSON.stringify(storedData));

        // 删除后刷新页面
        window.location.reload();
    } else {
        alert('Please select at least one row to delete.');
    }
});


const updateButton = document.querySelector('.update-button');
updateButton.addEventListener("click", function () {
    const checkedCheckboxes = document.querySelectorAll('input[type="checkbox"]:checked');
    const addressesToUpdate = []; // 用于存储选中行的 address 数据
    checkedCheckboxes.forEach(function (checkbox) {
        if (!checkbox.closest('th:first-child')) {
            const row = checkbox.closest('tr'); // 找到包含复选框的行
            const addressCell = row.querySelector('td:nth-child(3)'); // 假设 address 在第二列
            const address = addressCell.textContent.trim(); // 获取 address 数据并去除空白字符
            addressesToUpdate.push(address); // 存储到数组中
        }
    });
    const addressesToUpdateString = addressesToUpdate.join(' ');
    console.log(addressesToUpdateString);
    const requestUrl = '/update?address=' + encodeURIComponent(addressesToUpdateString);
    showLoadingModal();

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
        hideLoadingModal();
        const storedAccountsData = localStorage.getItem('accountsData');
        if (storedAccountsData) {
            // 解析已存储的数据
            let accountsData = JSON.parse(storedAccountsData);
            // 遍历新数据中的每个账户
            data.Accounts.forEach(function (newAccount) {
                console.log(data.Accounts[0].address)
                const existingAccountIndex = accountsData.findIndex(function (existingAccount) {
                    return existingAccount.address === newAccount.address;
                });
            
                if (existingAccountIndex !== -1) {
                    // 如果已经存在相同地址的账户，则用新数据替换旧数据
                    accountsData[existingAccountIndex] = newAccount;
                } else {
                    // 如果不存在相同地址的账户，则将新账户数据添加到数组中
                    accountsData.push(newAccount);
                    
                }
                console.log(newAccount)
            });
            
            console.log(data.Accounts)
            // 更新 localStorage 中的数据
            console.log(accountsData)
            localStorage.setItem('accountsData', JSON.stringify(accountsData));
        } else {
            // 如果之前没有存储的数据，直接存储新数据
            localStorage.setItem('accountsData', JSON.stringify(data.Accounts));
        }
        console.log(localStorage.getItem('accountsData'));
        // 渲染数据到表格
        window.location.reload();
        renderDataToTable(accountsData);
    })
    .catch(function (error) {
        console.error('Error:', error);
    });
     
    
});



document.querySelector('.address').addEventListener('submit', function (e) {
    e.preventDefault(); // 阻止默认的表单提交行为

    // 获取输入框中的地址
    var address = document.getElementById('address').value;

    // 构建请求URL，将用户输入的地址作为参数
    var requestUrl = '/BatchQuery?address=' + encodeURIComponent(address);

    // 发起AJAX请求

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
        
        // 获取之前存储在localStorage中的数据
        const storedAccountsData = localStorage.getItem('accountsData');
        if (storedAccountsData) {
            // 解析已存储的数据
            let accountsData = JSON.parse(storedAccountsData);
            // 遍历新数据中的每个账户
            data.Accounts.forEach(function (newAccount) {
                console.log(data.Accounts[0].address)
                const existingAccountIndex = accountsData.findIndex(function (existingAccount) {
                    return existingAccount.address === newAccount.address;
                });
            
                if (existingAccountIndex !== -1) {
                    // 如果已经存在相同地址的账户，则用新数据替换旧数据
                    accountsData[existingAccountIndex] = newAccount;
                } else {
                    // 如果不存在相同地址的账户，则将新账户数据添加到数组中
                    accountsData.push(newAccount);
                    
                }
                console.log(newAccount)
            });
            
            console.log(data.Accounts)
            // 更新 localStorage 中的数据
            console.log(accountsData)
            localStorage.setItem('accountsData', JSON.stringify(accountsData));
        } else {
            // 如果之前没有存储的数据，直接存储新数据
            localStorage.setItem('accountsData', JSON.stringify(data.Accounts));
        }
        console.log(localStorage.getItem('accountsData'));
        // 渲染数据到表格
        window.location.reload();
        renderDataToTable(accountsData);
        
    })
    .catch(function (error) {
        console.error('Error:', error);
    });
    document.getElementById('address').value = '';
});


const storedAccountsData = localStorage.getItem('accountsData');
    if (storedAccountsData) {
        // 将存储的数据解析为 JavaScript 对象
        const accountsData = JSON.parse(storedAccountsData);
        // 渲染数据到表格
        renderDataToTable({ Accounts: accountsData });
    }


    const selectAll = document.getElementById("selcetall");
    const checkboxes = document.querySelectorAll('input[type="checkbox"]');
    // console.log(selectAll);
    // console.log(checkboxes);

    
    // 全选/取消全选
    selectAll.addEventListener("change", function () { 
        checkboxes.forEach(function (checkbox) {
            checkbox.checked = selectAll.checked; 
        });
    });

    // 删除选中行
    // const deleteBtn = document.querySelector('.delete-button');
    // deleteBtn.addEventListener("click", function () {
    //     const checkedCheckboxes = document.querySelectorAll('input[type="checkbox"]:checked');
    //     checkedCheckboxes.forEach(function (checkbox) {
    //         const rowNumber = parseInt(checkbox.closest('tr').querySelector('td:first-child + td').textContent);
    //         if (!isNaN(rowNumber)) {
    //             // 删除表格中的行
    //             const rowToDelete = tbody.querySelector(`tr:nth-child(${rowNumber})`);
    //             if (rowToDelete) {
    //                 rowToDelete.remove();
    //             }
    //             const storedData = JSON.parse(localStorage.getItem('accountsData')) || [];
    //             storedData.splice(rowNumber - 1, 1);
    //             localStorage.setItem('accountsData', JSON.stringify(storedData));
    //         }
    //     });
    //     // 取消全选状态
    //     selectAll.checked = false;
    // });



    
});
