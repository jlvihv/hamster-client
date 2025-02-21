<template>
  <PageWrapper>
    <div id="terminal" class="xterm"></div>
    <Button @click="goBack" type="primary">{{ t('common.cancelText') }}</Button>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import 'xterm/css/xterm.css';
  import { Terminal } from 'xterm';
  import { FitAddon } from 'xterm-addon-fit';
  import { AttachAddon } from 'xterm-addon-attach';
  import { getCurrentInstance, onMounted } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button } from 'ant-design-vue';
  import { JudgeP2pReconnection } from '/@wails/go/app/P2p';
  import { ReconnectionProLink } from '/@wails/go/app/Application';
  import { useMessage } from '/@/hooks/web/useMessage';
  const { proxy } = getCurrentInstance();
  const { t } = useI18n();
  const router = useRouter();
  const { params } = useRoute();
  const applicationId = Number(params.id);
  const { createErrorModal } = useMessage();
  const goBack = () => {
    router.push('/applications/' + applicationId);
  };
  onMounted(() => {
    initTerm();
  });
  const initTerm = async () => {
    const isReconnection = await JudgeP2pReconnection();
    if (isReconnection) {
      const res = await ReconnectionProLink();
      if (!res) {
        createErrorModal({
          title: t('common.errorTip'),
          content: t('applications.cli.p2pReconnectionError'),
        });
        return;
      }
    }
    const term = new Terminal({});
    let socket = new WebSocket(`ws://localhost:10771/api/v1/thegraph/ws?serviceName=index-cli`);
    const attachAddon = new AttachAddon(socket);
    const fitAddon = new FitAddon();
    term.loadAddon(attachAddon);
    term.loadAddon(fitAddon);
    term.open(document.getElementById('terminal') as HTMLElement);
    fitAddon.fit();
    term.focus();
    proxy.term = term;
    proxy.socket = socket;
  };
</script>

<style scoped></style>
