
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
        <el-form-item label="机器人ID:" prop="botId">
    <el-input v-model.number="formData.botId" :clearable="true" placeholder="请输入机器人ID" />
</el-form-item>
        <el-form-item label="机器人类型:" prop="botType">
    <el-input v-model.number="formData.botType" :clearable="true" placeholder="请输入机器人类型" />
</el-form-item>
        <el-form-item label="创作者用户 ID:" prop="creatorUserId">
    <el-input v-model.number="formData.creatorUserId" :clearable="true" placeholder="请输入创作者用户 ID" />
</el-form-item>
        <el-form-item label="令牌:" prop="token">
    <el-input v-model="formData.token" :clearable="true" placeholder="请输入令牌" />
</el-form-item>
        <el-form-item label="描述:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="机器人聊天记录:" prop="botChatHistory">
    <el-switch v-model="formData.botChatHistory" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="机器人无聊天:" prop="botNochats">
    <el-switch v-model="formData.botNochats" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否验证:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="botInlineGeo字段:" prop="botInlineGeo">
    <el-switch v-model="formData.botInlineGeo" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="botInfoVersion字段:" prop="botInfoVersion">
    <el-input v-model.number="formData.botInfoVersion" :clearable="true" placeholder="请输入botInfoVersion字段" />
</el-form-item>
        <el-form-item label="botInlinePlaceholder字段:" prop="botInlinePlaceholder">
    <el-input v-model="formData.botInlinePlaceholder" :clearable="true" placeholder="请输入botInlinePlaceholder字段" />
</el-form-item>
        <el-form-item label="创建于:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="创建于:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTeamChaliBots,
  updateTeamChaliBots,
  findTeamChaliBots
} from '@/api/system/teamchalibots'

defineOptions({
    name: 'TeamChaliBotsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            botId: undefined,
            botType: undefined,
            creatorUserId: undefined,
            token: '',
            description: '',
            botChatHistory: false,
            botNochats: false,
            verified: false,
            botInlineGeo: false,
            botInfoVersion: undefined,
            botInlinePlaceholder: '',
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTeamChaliBots({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createTeamChaliBots(formData.value)
               break
             case 'update':
               res = await updateTeamChaliBots(formData.value)
               break
             default:
               res = await createTeamChaliBots(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
