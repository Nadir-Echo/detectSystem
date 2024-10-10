<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="90px">
      <el-form-item label="用户名">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="姓名">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="负责生产线">
        <el-input v-model="form.proID" />
      </el-form-item>
      <el-form-item label="电话">
        <el-input v-model="form.phone" />
      </el-form-item>

      <el-form-item label="角色">
        <el-radio-group v-model="form.role">
          <el-radio label="admin" />
          <el-radio label="detector" />
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getToken } from '@/utils/auth'
import { addUser } from '@/api/manage'
export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
        name: '',
        proID: null,
        phone: '',
        role: ''
      }
    }
  },
  created() {
    this.authadmin()
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
    onSubmit() {
      //   this.$message('submit!')
      console.log(this.form)
      addUser(this.form).then(res => {
        if (res.code === 200) {
          this.$message({
            message: '添加成功！',
            type: 'success'
          })
        } else if (res.code === 50012) {
          this.$message({
            message: '用户名已存在！',
            type: 'error'
          })
        } else {
          this.$message({
            message: '添加失败!',
            type: 'error'
          })
        }
      })
    },
    onCancel() {
      this.$message({
        message: '取消!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>

