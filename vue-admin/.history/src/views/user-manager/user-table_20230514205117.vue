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
          {{ scope.row.userID }}
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

import { getUserInfo } from '@/api/user-manage'
export default {
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchUserData()
  },
  methods: {
    fetchUserData() {
      this.listLoading = true
      getUserInfo().then(response => {
        this.list = response.data
        this.listLoading = false
        console.log(this.list)
      })
    }
  }
}
</script>
