
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入用户ID" />
</el-form-item>
        <el-form-item label="userType字段:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入userType字段" />
</el-form-item>
        <el-form-item label="accessHash字段:" prop="accessHash">
    <el-input v-model.number="formData.accessHash" :clearable="true" placeholder="请输入accessHash字段" />
</el-form-item>
        <el-form-item label="secretKey用户ID:" prop="secretKeyId">
    <el-input v-model.number="formData.secretKeyId" :clearable="true" placeholder="请输入secretKey用户ID" />
</el-form-item>
        <el-form-item label="firstName字段:" prop="firstName">
    <el-input v-model="formData.firstName" :clearable="true" placeholder="请输入firstName字段" />
</el-form-item>
        <el-form-item label="lastName字段:" prop="lastName">
    <el-input v-model="formData.lastName" :clearable="true" placeholder="请输入lastName字段" />
</el-form-item>
        <el-form-item label="username字段:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入username字段" />
</el-form-item>
        <el-form-item label="phone字段:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入phone字段" />
</el-form-item>
        <el-form-item label="countryCode字段:" prop="countryCode">
    <el-input v-model="formData.countryCode" :clearable="true" placeholder="请输入countryCode字段" />
</el-form-item>
        <el-form-item label="verified字段:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="support字段:" prop="support">
    <el-switch v-model="formData.support" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="scam字段:" prop="scam">
    <el-switch v-model="formData.scam" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="fake字段:" prop="fake">
    <el-switch v-model="formData.fake" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="premium字段:" prop="premium">
    <el-switch v-model="formData.premium" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="premiumExpireDate字段:" prop="premiumExpireDate">
    <el-input v-model.number="formData.premiumExpireDate" :clearable="true" placeholder="请输入premiumExpireDate字段" />
</el-form-item>
        <el-form-item label="about字段:" prop="about">
    <el-input v-model="formData.about" :clearable="true" placeholder="请输入about字段" />
</el-form-item>
        <el-form-item label="state字段:" prop="state">
    <el-input v-model.number="formData.state" :clearable="true" placeholder="请输入state字段" />
</el-form-item>
        <el-form-item label="isBot字段:" prop="isBot">
    <el-switch v-model="formData.isBot" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="accountDaysTtl字段:" prop="accountDaysTtl">
    <el-input v-model.number="formData.accountDaysTtl" :clearable="true" placeholder="请输入accountDaysTtl字段" />
</el-form-item>
        <el-form-item label="photo用户ID:" prop="photoId">
    <el-input v-model.number="formData.photoId" :clearable="true" placeholder="请输入photo用户ID" />
</el-form-item>
        <el-form-item label="restricted字段:" prop="restricted">
    <el-switch v-model="formData.restricted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="restrictionReason字段:" prop="restrictionReason">
    <el-input v-model="formData.restrictionReason" :clearable="true" placeholder="请输入restrictionReason字段" />
</el-form-item>
        <el-form-item label="archiveAndMuteNewNoncontactPeers字段:" prop="archiveAndMuteNewNoncontactPeers">
    <el-switch v-model="formData.archiveAndMuteNewNoncontactPeers" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="emojiStatusDocument用户ID:" prop="emojiStatusDocumentId">
    <el-input v-model.number="formData.emojiStatusDocumentId" :clearable="true" placeholder="请输入emojiStatusDocument用户ID" />
</el-form-item>
        <el-form-item label="emojiStatusUntil字段:" prop="emojiStatusUntil">
    <el-input v-model.number="formData.emojiStatusUntil" :clearable="true" placeholder="请输入emojiStatusUntil字段" />
</el-form-item>
        <el-form-item label="storiesMax用户ID:" prop="storiesMaxId">
    <el-input v-model.number="formData.storiesMaxId" :clearable="true" placeholder="请输入storiesMax用户ID" />
</el-form-item>
        <el-form-item label="color字段:" prop="color">
    <el-input v-model.number="formData.color" :clearable="true" placeholder="请输入color字段" />
</el-form-item>
        <el-form-item label="colorBackgroundEmoji用户ID:" prop="colorBackgroundEmojiId">
    <el-input v-model.number="formData.colorBackgroundEmojiId" :clearable="true" placeholder="请输入colorBackgroundEmoji用户ID" />
</el-form-item>
        <el-form-item label="profileColor字段:" prop="profileColor">
    <el-input v-model.number="formData.profileColor" :clearable="true" placeholder="请输入profileColor字段" />
</el-form-item>
        <el-form-item label="profileColorBackgroundEmoji用户ID:" prop="profileColorBackgroundEmojiId">
    <el-input v-model.number="formData.profileColorBackgroundEmojiId" :clearable="true" placeholder="请输入profileColorBackgroundEmoji用户ID" />
</el-form-item>
        <el-form-item label="birthday字段:" prop="birthday">
    <el-input v-model="formData.birthday" :clearable="true" placeholder="请输入birthday字段" />
</el-form-item>
        <el-form-item label="personalChannel用户ID:" prop="personalChannelId">
    <el-input v-model.number="formData.personalChannelId" :clearable="true" placeholder="请输入personalChannel用户ID" />
</el-form-item>
        <el-form-item label="authorizationTtlDays字段:" prop="authorizationTtlDays">
    <el-input v-model.number="formData.authorizationTtlDays" :clearable="true" placeholder="请输入authorizationTtlDays字段" />
</el-form-item>
        <el-form-item label="deleted字段:" prop="deleted">
    <el-switch v-model="formData.deleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="deleteReason字段:" prop="deleteReason">
    <el-input v-model="formData.deleteReason" :clearable="true" placeholder="请输入deleteReason字段" />
</el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
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
  createUsers,
  updateUsers,
  findUsers
} from '@/api/system/users'

defineOptions({
    name: 'UsersForm'
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
            userType: undefined,
            accessHash: undefined,
            secretKeyId: undefined,
            firstName: '',
            lastName: '',
            username: '',
            phone: '',
            countryCode: '',
            verified: false,
            support: false,
            scam: false,
            fake: false,
            premium: false,
            premiumExpireDate: undefined,
            about: '',
            state: undefined,
            isBot: false,
            accountDaysTtl: undefined,
            photoId: undefined,
            restricted: false,
            restrictionReason: '',
            archiveAndMuteNewNoncontactPeers: false,
            emojiStatusDocumentId: undefined,
            emojiStatusUntil: undefined,
            storiesMaxId: undefined,
            color: undefined,
            colorBackgroundEmojiId: undefined,
            profileColor: undefined,
            profileColorBackgroundEmojiId: undefined,
            birthday: '',
            personalChannelId: undefined,
            authorizationTtlDays: undefined,
            deleted: false,
            deleteReason: '',
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
      const res = await findUsers({ ID: route.query.id })
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
               res = await createUsers(formData.value)
               break
             case 'update':
               res = await updateUsers(formData.value)
               break
             default:
               res = await createUsers(formData.value)
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
