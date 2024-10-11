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
      <el-table-column align="center" label="PID" width="95">
        <template slot-scope="scope">
          {{ scope.row.pid }}
        </template>
      </el-table-column>
      <el-table-column label="缺陷" align="center">
        <template slot-scope="scope">
          {{ scope.row.class }}
        </template>
      </el-table-column>
      <el-table-column label="生产线" width="110" align="center">
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
      <el-table-column align="center" label="Detail">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleDetail(scope.$index, scope.row)"
          >查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>