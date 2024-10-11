<template>
  <div class="app-container">
    <el-table>
      <el-table-column align="center" label="PID" width="95">
        <template slot-scope="scope">
          {{ scope.row.userId }}
        </template>
      </el-table-column>
      <el-table-column label="缺陷" align="center">
        <template slot-scope="scope">
          {{ scope.row.userName }}
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
        label="Defect_time"
        width="200"
      >
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.time }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
