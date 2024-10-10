<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="PID" width="85">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          {{ scope.row.user_name }}
        </template>
      </el-table-column>

      <el-table-column label="姓名" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="负责生产线" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pro_id }}</span>
        </template>
      </el-table-column>

      <el-table-column
        align="center"
        prop="created_at"
        label="角色"
        width="200"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.role }}</span>
        </template>
      </el-table-column>
      <el-table-column label="电话" align="center">
        <template slot-scope="scope">
          {{ scope.row.phone }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

import { getUserInfo } from '@/api/manage'
import { getToken } from '@/utils/auth'
export default {
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.authadmin(),
    this.fetchUserData()
  },
  methods: {
    authadmin() {
      const myHeaders = new Headers()
      myHeaders.append('X-Token', getToken())
      fetch('http://localhost:8010/auth/admin', {
        method: 'GET',
        headers: myHeaders
      })
        .then((response) => response.json()) // 解析响应数据为JSON对象
        .then((jsonResponse) => {
          // judge jsonResponse.data is 50009
          if (jsonResponse.code === 50009) {
            this.$message({
              message: '用户权限不足',
              type: 'warning'
            })
            // page to 404
            this.$router.push({ path: '/404' })
          } else {
            // 从JSON对象中获取需要的数据
            const code = jsonResponse.code
            // const data = jsonResponse.data

            // 在控制台输出获取的数据，用于测试
            console.log(`code: ${code}`)
            // console.log(`data: ${data}`)
          }
        })
        .catch((error) => {
          console.log(error)
        })
    },
    fetchUserData() {
      this.listLoading = true
      getUserInfo().then(response => {
        this.list = response.data
        this.listLoading = false
        // console.log(this.list)
      })
    }
  }
}
</script>
