<template>
  <div class="app-container">
    <el-table>
      <el-table-column align="center" label="PID" width="85">
        <template slot-scope="scope">
          {{ scope.row.userId }}
        </template>
      </el-table-column>
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          {{ scope.row.userName }}
        </template>
      </el-table-column>

      <el-table-column label="姓名" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="负责生产线" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.proId }}</span>
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
imp