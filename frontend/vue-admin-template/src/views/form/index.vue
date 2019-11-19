<template>
  <div class="app-container">
    <el-form ref="dataForm" :model="form" label-width="120px">
      <el-form-item label="域名记录">
        <el-input v-model="form.domain" />
      </el-form-item>
      <el-form-item label="解析类型">
        <el-select v-model="form.region" placeholder="please select your zone">
          <el-option label="A记录" value="A" />
        </el-select>
      </el-form-item>
      <el-form-item label="解析时间">
        <el-input v-model="form.ttl" />
      </el-form-item>
      <el-form-item label="IPv4地址">
        <el-input v-model="form.host" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">创建</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { addDomain } from '@/api/table'

export default {
  data() {
    return {
      form: {
        domain: '',
        host: '',
        ttl: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$refs.dataForm.validate((valid) => {
        if (valid) {
          addDomain(this.form).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: '添加成功',
              type: 'warning'
            })
          })
        }
      })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

