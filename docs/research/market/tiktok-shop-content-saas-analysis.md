# Plano de Produto e Arquitetura: SaaS de IA para TikTok Shop, Afiliados e Conteudo Curto

Data da pesquisa: 2026-07-01

## 1. Tese do Produto

A oportunidade nao e criar uma ferramenta de spam, scraping ou automacao cega. A oportunidade defensavel e construir um sistema operacional de conteudo curto para criadores, afiliados, agencias e pequenas marcas que precisam transformar produtos, tendencias e videos longos em roteiros, cortes, publicacoes e analises com velocidade, mas usando APIs oficiais, autorizacao OAuth, direitos de uso, disclosure de IA e limites de plataforma.

Posicionamento recomendado:

- "AI content operations for TikTok Shop affiliates and short-form commerce."
- Foco inicial: TikTok Shop + TikTok video workflow.
- Expansao: YouTube Shorts, Instagram Reels, Facebook Reels e agendamento multi-canal.
- Diferencial: nao competir apenas com editores como CapCut/OpusClip; competir com uma camada de produto comercial: pesquisa de produto, biblioteca de hooks/CTAs, compliance, creditos de IA, pipeline de cortes, analytics e aprendizado de performance por nicho/produto.

Principios inegociaveis:

- Somente publicacao por API oficial ou fluxo manual assistido.
- Nunca automatizar login, navegador, curtidas, comentarios, follows, views ou engajamento.
- Sempre manter consentimento do usuario, OAuth e logs de acao.
- Detectar e avisar sobre conteudo reutilizado, marcas d'agua, material sem direitos, claims medicos/financeiros e AIGC que exige disclosure.
- Em conteudo de cortes, exigir transformacao substantiva: comentario, narrativa, edicao relevante, contexto, fonte e direitos.

## 2. Mercado e Plataformas

### 2.1 TikTok Shop e TikTok

O TikTok Shop combina descoberta, video curto, live commerce, afiliados e checkout in-app. O produto SaaS deve tratar TikTok como o canal inicial porque a dor e mais especifica: afiliados precisam encontrar produtos, criar criativos, testar angulos, postar com consistencia e entender conversao.

APIs e ferramentas oficiais relevantes:

- TikTok for Developers lista Login Kit, Share Kit, Content Posting API, Display API, Research Tools, Commercial Content API, Data Portability API e recursos de monetizacao ([TikTok Developers](https://developers.tiktok.com/doc/content-posting-api-get-started)).
- Content Posting API permite postagem direta de video/foto, mas exige app registrado, produto habilitado, escopo aprovado `video.publish`, autorizacao do usuario e auditoria para remover restricao de visibilidade privada em clientes nao auditados ([Content Posting API](https://developers.tiktok.com/doc/content-posting-api-get-started)).
- Query Creator Info retorna configuracoes como privacidade, duet, stitch e `max_video_post_duration_sec`, que devem ser checadas antes de publicar ([Content Posting API](https://developers.tiktok.com/doc/content-posting-api-get-started)).
- Rate limits padrao para endpoints de Display API sao 600 requisicoes por minuto por endpoint em janela deslizante; exceder retorna HTTP 429 `rate_limit_exceeded` ([Rate Limits](https://developers.tiktok.com/doc/tiktok-api-v2-rate-limit)).
- Research Tools incluem endpoints de TikTok Shop info/products/reviews, mas esse caminho tende a ser voltado a pesquisa aprovada e nao deve ser assumido como API comercial livre para scraping de mercado ([TikTok Developers nav](https://developers.tiktok.com/doc/content-posting-api-get-started)).
- TikTok Creative Center/TikTok One e ferramentas oficiais de criacao sao fontes importantes para tendencias e inspiracao, mas a automacao deve respeitar os termos e nao raspar conteudo sem permissao ([TikTok Creative Center](https://ads.tiktok.com/creative/creativeCenter)).

Politicas criticas:

- TikTok exige disclosure para conteudo gerado por IA ou significativamente editado quando mostrar pessoas/cenas realistas; conteudo nao rotulado pode ser removido, restringido ou rotulado pela plataforma ([TikTok Community Guidelines](https://www.tiktok.com/community-guidelines/en/integrity-authenticity/)).
- TikTok nao permite AIGC enganoso sobre temas de importancia publica, uso de likeness de figuras privadas sem consentimento, imitacao de voz real sem disclosure adequado e outros cenarios sensiveis ([TikTok Community Guidelines](https://www.tiktok.com/community-guidelines/en/integrity-authenticity/)).
- Conteudo reutilizado/unoriginal sem algo novo fica inelegivel para For You feed; exemplos incluem clipes com marca d'agua/logo de terceiros, GIFs ou edicoes minimas ([TikTok Community Guidelines](https://www.tiktok.com/community-guidelines/en/integrity-authenticity/)).
- TikTok proibe ferramentas/scripts/truques para burlar sistemas, manipular engajamento ou inflar metricas; isso pode gerar remocao, restricao ou banimento ([TikTok Community Guidelines](https://www.tiktok.com/community-guidelines/en/integrity-authenticity/)).

Oportunidades:

- "Compliance by design": checklist automatico antes de exportar/postar.
- Pipeline de afiliados: produto -> angulos -> scripts -> video -> caption -> hashtags -> publicacao -> analytics -> iteracao.
- Score interno de risco: AIGC disclosure, marcas d'agua, conteudo duplicado, claims proibidos, direitos de musica/imagem.
- Product research oficial ou semi-manual: integrar Seller/Creator workflows apenas quando houver permissao/API; caso contrario, criar "research notebook" onde o usuario cola links/dados e o sistema estrutura.
- Diferenciacao forte em TikTok Shop: concorrentes de video fazem cortes, mas poucos conectam diretamente produto, criativo, afiliado e conversao.

### 2.2 Programa de Afiliados e Creator Marketplace

O programa de afiliados conecta criadores e sellers por comissao. A plataforma deve tratar afiliacao como workflow comercial:

- cadastro de produtos/campanhas;
- comissoes e categorias;
- amostras recebidas;
- status de roteiro/gravação/postagem;
- links/codigos;
- performance por produto, criativo e criador.

Lacuna tecnica: TikTok Shop Partner/Open APIs existem em portal proprio, mas o acesso e permissao variam por regiao, app type e partner approval. Portanto, o MVP nao deve depender de acesso irrestrito a todos os dados de afiliado. O MVP deve suportar importacao manual/CSV e links de produto, e deixar conectores oficiais como fase 1.0/2.0.

### 2.3 Instagram Reels e Facebook Reels

O Instagram e Facebook devem entrar como canais de distribuicao e analytics, nao como fonte primaria de scraping.

APIs e restricoes:

- O caminho oficial e Meta Graph API/Instagram Platform para contas profissionais, com OAuth, permissoes, content publishing e insights. A documentacao oficial de Meta deve ser validada durante a implementacao porque endpoints e permissoes mudam com frequencia ([Meta Developers](https://developers.facebook.com/docs/)).
- Meta Community Standards proibe comportamento inautentico, redes de ativos falsos e tentativa de enganar Meta ou usuarios, o que reforca que o SaaS nao deve automatizar contas ou engajamento ([Meta Transparency Center](https://transparency.meta.com/policies/community-standards/inauthentic-behavior/)).
- Para Facebook Pages/Reels, a publicacao passa por Graph API/Page video publishing e permissoes de pagina; na pratica, exigir app review e tokens de longa duracao com refresh controlado.

Oportunidades:

- Exportar formatos Reels prontos com safe areas, capas, titulo, legenda, hashtags e arquivos de legenda.
- Agendamento oficial onde permitido; quando nao permitido, gerar pacote de publicacao manual.
- Analytics comparativo TikTok/Reels/Shorts por mesmo corte, hook, CTA e produto.

### 2.4 YouTube Shorts

APIs e restricoes:

- YouTube Data API `videos.insert` faz upload e define metadata; projetos nao verificados criados apos 2020 ficam com videos privados ate auditoria de compliance ([YouTube videos.insert](https://developers.google.com/youtube/v3/docs/videos/insert)).
- Upload aceita video ate 256 GB, requer OAuth com escopos como `youtube.upload`, e custa quota de Video Uploads; a doc informa impacto de 100 chamadas/dia no bucket de upload ([YouTube videos.insert](https://developers.google.com/youtube/v3/docs/videos/insert)).
- YouTube exige originalidade/autenticidade para monetizacao. Conteudo massificado, repetitivo ou gerado por IA com template generico e sem perspectiva original pode ser inelegivel ([YouTube Monetization Policies](https://support.google.com/youtube/answer/1311392)).
- Conteudo reutilizado precisa de diferenca significativa, comentario original, modificacao substantiva ou valor educativo/entretenimento; permissao do autor nao garante monetizacao se o conteudo continuar reutilizado sem transformacao ([YouTube Monetization Policies](https://support.google.com/youtube/answer/1311392)).
- YouTube tem disclosure para conteudo gerado/alterado por IA em certos casos ([YouTube GenAI disclosure](https://support.google.com/youtube/answer/14328491)).

Oportunidades:

- Shorts como canal de diversificacao para cortes.
- Modulo de "transformacao suficiente" com checklist: narracao original, contexto, efeitos visuais, capitulo, comentario e fontes.
- Metadata generator com campo `status.containsSyntheticMedia` quando aplicavel.

## 3. Concorrentes

### 3.1 Opus Clip

Funcionalidades: AI clipping, virality score, captions, AI reframe, B-roll, scheduler, brand templates, XML export, team workspace e API em planos business ([OpusClip pricing](https://www.opus.pro/pricing)).

Pontos fortes:

- Muito forte em long video -> shorts.
- Virality score claro para criadores.
- Editor + captions + reframe + scheduler em fluxo unico.
- Posicionamento forte para creators, podcasters, agencias e e-commerce.

Pontos fracos/oportunidades:

- Nao e centrado em TikTok Shop affiliate commerce.
- Pesquisa de produtos, comissoes, criativos por produto e conversao nao parecem ser o core.
- Compliance de AIGC/reused content/direitos poderia ser mais explicito como produto.

Modelo: freemium + assinatura por creditos + business custom.

### 3.2 Captions

Funcionalidades: edicao mobile/desktop, captions, AI creator/avatar/video tools, geracao e ajustes de videos.

Pontos fortes:

- Produto muito orientado a criador individual.
- Excelente UX para legendas e video rapido.
- Forte em AI presentation/avatar e conteudo selfie.

Pontos fracos/oportunidades:

- Menos orientado a operacao comercial de afiliados.
- Analytics, pesquisa de produto e pipeline multi-perfil nao sao seu centro.

Modelo: assinatura com planos de creator/pro.

Fonte: [Captions pricing](https://captions.ai/pricing).

### 3.3 Descript

Funcionalidades: edicao por texto, transcricao, captions, studio sound, overdub/voz, templates e colaboracao.

Pontos fortes:

- Melhor categoria para edicao baseada em texto.
- Excelente para podcast/video longo e equipes.
- Fluxo de transcricao maduro.

Pontos fracos/oportunidades:

- Nao e verticalizado para TikTok Shop.
- Menos prescritivo em hooks, produtos, afiliados e conversao.

Modelo: freemium + subscriptions por usuario/uso.

Fonte: [Descript pricing](https://www.descript.com/pricing).

### 3.4 Submagic

Funcionalidades: captions, B-roll, hooks, zooms, efeitos, shorts e templates.

Pontos fortes:

- Otimo em "legendas que parecem TikTok".
- Foco forte em creator economy.

Pontos fracos/oportunidades:

- Menos forte em gestao de campanhas/produtos/analytics comercial.
- Oportunidade para competir com compliance + TikTok Shop + workflow de afiliados.

Modelo: assinatura mensal/anual por volume.

Fonte: [Submagic pricing](https://www.submagic.co/pricing).

### 3.5 CapCut

Funcionalidades: editor desktop/mobile/online, efeitos, templates, auto captions, text-to-speech, AI voice, AI dubbing, long video to shorts, upscaling, script-to-video, busca inteligente e compartilhamento para TikTok/YouTube ([CapCut](https://www.capcut.com/tools/desktop-video-editor)).

Pontos fortes:

- Distribuicao massiva e proximidade com ecossistema ByteDance/TikTok.
- Editor poderoso e familiar para criadores.
- Muitas ferramentas AI e templates.

Pontos fracos/oportunidades:

- Nao e um SaaS B2B de operacao de afiliados.
- Pode ser ferramenta integrada/export target, nao necessariamente concorrente frontal.
- Falta camada profunda de produto/comissao/performance por criativo.

Modelo: freemium/pro.

### 3.6 Metricool

Funcionalidades: social planner, analytics, reports, competitor analysis, inbox, ads, AI assistant, multiplataforma e Metricool MCP ([Metricool pricing](https://metricool.com/pricing/)).

Pontos fortes:

- Forte em planejamento e analytics.
- Bom para agencias e marcas.
- Cobre muitos canais.

Pontos fracos/oportunidades:

- Nao e editor de video/IA de cortes profundo.
- Nao e verticalizado em TikTok Shop affiliates.

Modelo: freemium + planos por numero de marcas.

### 3.7 Buffer

Funcionalidades: agendamento, publicacao, calendario, analytics e colaboracao social.

Pontos fortes:

- Simplicidade, confianca, preco acessivel.
- Bom para pequenas empresas.

Pontos fracos/oportunidades:

- IA/video commerce nao e centro.
- Nao resolve criacao de cortes, scripts e TikTok Shop.

Modelo: freemium + assinatura por canal/usuario.

Fonte: [Buffer pricing](https://buffer.com/pricing).

### 3.8 Hootsuite

Funcionalidades: gestao social enterprise, agendamento, inbox, analytics, governanca e times.

Pontos fortes:

- Enterprise, compliance, permissoes e workflow de aprovacao.
- Forte para agencias e grandes marcas.

Pontos fracos/oportunidades:

- Caro/complexo para criadores e afiliados.
- Pouco vertical para IA de video curto e TikTok Shop.

Modelo: assinatura premium/enterprise.

Fonte: [Hootsuite plans](https://www.hootsuite.com/plans).

### 3.9 Repurpose.io

Funcionalidades: republicacao/distribuicao entre TikTok, YouTube Shorts, Facebook, Instagram e outros; planos por contas conectadas e volume ([Repurpose.io pricing](https://repurpose.io/pricing/)).

Pontos fortes:

- Automacao de distribuicao muito clara.
- Bom para criadores com muitos canais.

Pontos fracos/oportunidades:

- Menos IA criativa e menos orientado a produto/afiliacao.
- Compliance e transformacao de conteudo podem ser diferencial do novo SaaS.

Modelo: assinatura por plano Starter/Pro/Agency.

### 3.10 vidIQ

Funcionalidades: pesquisa de keywords, ideacao, SEO, analytics e crescimento para YouTube.

Pontos fortes:

- YouTube SEO e ideacao.
- Forte em creators que querem crescer canal.

Pontos fracos/oportunidades:

- Nao e editor nem TikTok Shop.
- Shorts commerce e afiliados sao lacunas.

Modelo: freemium + assinatura.

Fonte: [vidIQ](https://vidiq.com/).

### 3.11 TubeBuddy

Funcionalidades: keyword explorer, SEO Studio, A/B testing, thumbnail analyzer/generator, chapter editor, title generator e insights ([TubeBuddy pricing](https://www.tubebuddy.com/pricing)).

Pontos fortes:

- YouTube SEO e produtividade.
- Extensao/integração madura com YouTube.

Pontos fracos/oportunidades:

- Nao resolve TikTok Shop, Reels, edicao de video ou afiliados.

Modelo: freemium + assinatura por plano.

## 4. Personas e Dores

Afiliado TikTok Shop iniciante:

- Nao sabe escolher produto nem angulo.
- Demora para criar roteiro, hook, CTA e legenda.
- Tem medo de banimento ou conteudo derrubado.
- Nao sabe medir qual video vendeu.
- Precisa de templates simples e baixo custo.

Criador iniciante:

- Nao domina edicao.
- Nao sabe manter calendario.
- Precisa transformar ideias em posts prontos.
- Busca consistencia e feedback rapido.

Pagina de cortes:

- Precisa importar videos longos, detectar momentos, cortar, legendar e postar em volume.
- Risco alto de conteudo reutilizado sem valor.
- Precisa provar transformacao: comentario, narrativa, contexto, edicao e direitos.

Agencia:

- Gerencia varias marcas/perfis.
- Precisa de aprovacao, permissoes, brand kits, relatorios e SLA.
- Quer reduzir custo de edicao e aumentar volume.

Pequena empresa/e-commerce:

- Tem produto, mas nao tem motor de conteudo.
- Precisa gerar criativos por SKU.
- Quer publicar em TikTok/Instagram/YouTube sem contratar equipe grande.

Influencer:

- Precisa preservar voz e marca pessoal.
- Quer monetizar sem parecer anuncio ruim.
- Precisa organizar demandas, produtos recebidos, briefings e prazos.

## 5. Funcionalidades Modulares

### 5.1 Dashboard

- Metricas: posts, views, CTR, engajamento, vendas atribuiveis, receita estimada, custo IA, taxa de aprovacao.
- Pipeline: Ideia -> Roteiro -> Gravacao -> Edicao -> Revisao -> Agendado -> Publicado -> Analisado.
- Fila: jobs de transcricao, render, geracao IA, export e publicacao.
- Agenda: calendario multi-canal.
- Projetos: por nicho, cliente, produto, canal e campanha.

API oficial: parcial. Analytics/publicacao dependem de TikTok/YouTube/Meta OAuth e permissoes. Pipeline interno sem custo externo direto.

### 5.2 Pesquisa

- Buscar produtos: no MVP por cadastro manual/CSV/link; depois TikTok Shop APIs oficiais se aprovadas.
- Buscar tendencias/hashtags/musicas: usar fontes oficiais quando disponiveis; evitar scraping proibido.
- Buscar concorrentes/videos virais: somente com dados publicos autorizados, APIs oficiais, embed ou entrada manual do usuario.
- Criar "research cards": produto, publico, promessa, prova, restricoes, exemplos permitidos.

Limitacao: APIs de tendencia/musica sao restritas. Nao prometer scraping universal.

### 5.3 IA

- Gerar roteiro por produto/nicho.
- Reescrever roteiro por tom.
- Criar 5-20 versoes de hook.
- Gerar CTA, titulos, descricoes e hashtags.
- Criar calendario.
- Classificar risco de policy e claims.
- Gerar sugestoes baseadas em analytics.

API: OpenAI/Anthropic/Gemini via SDKs oficiais; custo por token/minuto/imagem/audio. Rate limit depende de conta/provedor.

### 5.4 Videos

- Upload para S3.
- Organizacao por projeto/produto/nicho.
- Biblioteca, versoes e assets.
- Detecao de cenas com ffmpeg + modelos vision.
- Transcricao com Whisper/OpenAI/Gemini/servicos ASR.
- Legendas SRT/VTT/ASS.
- Traducao e dublagem.
- Auto zoom/reframe.
- Remocao de silencio.
- Highlight detection.

API: grande parte e interna usando ffmpeg + IA. Custo vem de armazenamento, processamento CPU/GPU e APIs de transcricao/vision/TTS.

### 5.5 Edicao

- Templates visuais.
- Aplicar templates em lote.
- Inserir legendas, emojis, efeitos, B-roll e musica licenciada.
- Exportar 9:16, 1:1, 16:9.
- Safe areas por plataforma.

Implementacao: ffmpeg/MoviePy/Remotion. Para escalar, jobs assíncronos e workers separados. Musica deve vir de biblioteca licenciada ou upload do usuario com confirmacao de direitos.

### 5.6 Conteudo

- Gerenciar paginas/perfis.
- Organizar nichos.
- Biblioteca de prompts, CTAs, hooks, thumbnails e brand voice.
- Banco de produtos e campanhas.
- Biblioteca de claims permitidos/proibidos por nicho.

### 5.7 Analytics

- Views, CTR, likes, comments, shares, watch time quando API permitir.
- Conversao e produtos vendidos via importacao CSV, pixel, UTMs, affiliate reports ou API oficial.
- Comparacao entre videos.
- Sugestoes IA: "hook A performou melhor em beleza", "videos com prova social convertem mais".

Limitacao: atribuicao de venda em TikTok Shop pode exigir acesso oficial/relatorios do seller/creator. MVP deve aceitar CSV/manual.

## 6. Modulo de Paginas de Cortes

Politicas:

- TikTok: conteudo reutilizado sem algo novo e ineligible para FYF; clipes com marca d'agua/logo de terceiros e edicoes minimas sao risco ([TikTok Guidelines](https://www.tiktok.com/community-guidelines/en/integrity-authenticity/)).
- YouTube: reused content precisa comentario, modificacao substantiva ou valor novo; permissao do criador nao elimina risco de monetizacao ([YouTube Monetization Policies](https://support.google.com/youtube/answer/1311392)).
- Meta: comportamento inautentico e ativos falsos sao proibidos; publicar em escala com identidades enganosas e alto risco ([Meta Transparency Center](https://transparency.meta.com/policies/community-standards/inauthentic-behavior/)).

Pipeline permitido:

1. Importar video com confirmacao de direitos: proprio, licenciado, permissao, dominio publico ou uso transformativo avaliado.
2. Extrair audio, transcrever e diarizar speakers.
3. Detectar assuntos, mudancas de cena, energia, perguntas, risadas, demonstracoes e momentos de tensao.
4. Gerar candidatos de corte com motivo: hook, payoff, pergunta, prova, controversia, tutorial.
5. Exigir camada de valor: comentario do criador, contexto textual, narracao, zoom/reframe, captions originais, B-roll proprio, fontes, opiniao ou resumo.
6. Gerar titulo, legenda, hashtags e capitulos.
7. Rodar checker: marcas d'agua, duplicidade, AIGC disclosure, direitos, claims.
8. Exportar Shorts/Reels/TikTok.
9. Publicar por API oficial ou pacote manual.

Tecnicas permitidas para agregar valor:

- Comentario original em audio/video.
- Contextualizacao: "por que isso importa".
- Analise de produto/mercado.
- Edicao substantiva com narrativa propria.
- Comparacao, tutorial, review, antes/depois.
- Credito/fonte quando aplicavel.
- Evitar simples compilacao ou repost.

## 7. Arquitetura Recomendada

### 7.1 Stack

- Backend: Golang.
- API: REST inicialmente; GraphQL apenas se o frontend ficar muito complexo.
- Banco: PostgreSQL.
- Fila: Redis Streams no MVP pela simplicidade; RabbitMQ se precisar roteamento complexo, retries dead-letter e multiplos tipos de workers.
- Storage: S3 compativel.
- Frontend: Next.js + React.
- Render/Media workers: Go orquestrando ffmpeg; Remotion/Node se templates forem React-based.
- Deploy: Docker Compose no MVP; Kubernetes quando houver varios workers, autoscaling e tenants enterprise.

### 7.2 Servicos

- `api-gateway`: auth, tenants, billing, REST.
- `identity-service`: usuarios, workspace, roles, OAuth tokens criptografados.
- `content-service`: projetos, roteiros, hooks, CTAs, prompts, calendario.
- `media-service`: uploads, transcricao, assets, render jobs.
- `ai-service`: provider abstraction para OpenAI/Anthropic/Gemini/local.
- `publishing-service`: TikTok/YouTube/Meta connectors e publicacao manual package.
- `analytics-service`: ingestao de metricas e vendas.
- `compliance-service`: regras de policy, direitos, AIGC, duplicate/reused risk.
- `billing-service`: planos, creditos, usage ledger.
- `worker-media`: ffmpeg, scenes, captions, exports.
- `worker-ai`: batch de scripts, titles, hashtags, classifiers.
- `worker-publish`: agendamento e retry.

### 7.3 Modelo de Dados Inicial

- `workspaces`, `users`, `memberships`, `roles`.
- `social_accounts`, `oauth_tokens`, `publishing_profiles`.
- `products`, `affiliate_campaigns`, `commission_rules`.
- `projects`, `content_items`, `scripts`, `script_versions`.
- `media_assets`, `media_versions`, `transcripts`, `captions`.
- `render_jobs`, `ai_jobs`, `publish_jobs`, `job_events`.
- `prompt_templates`, `hook_bank`, `cta_bank`, `thumbnail_bank`.
- `analytics_snapshots`, `sales_imports`, `video_performance`.
- `compliance_checks`, `policy_findings`, `rights_confirmations`.
- `usage_ledger`, `credit_transactions`, `plans`.

### 7.4 Seguranca e Compliance

- Tokens OAuth criptografados com KMS/Secrets Manager.
- Auditoria de toda publicacao e export.
- Rate limiting por workspace e provider.
- Idempotency keys para jobs/publicacoes.
- Isolamento tenant via `workspace_id` em todas as tabelas; RLS no PostgreSQL se necessario.
- LGPD/GDPR: export/delete data, consentimento, data retention.
- Moderacao de prompts e outputs para claims proibidos.
- Logs sem dados sensiveis de token/audio cru.

## 8. Modelos de IA

### 8.1 OpenAI

Uso recomendado:

- GPT-5.4 mini/nano ou similar de baixo custo: titulos, hashtags, CTAs, classificacao simples.
- GPT-5.5/GPT-5.4: roteiros, analise estrategica, prompts complexos, avaliacao de compliance.
- Realtime/Whisper/transcription: transcricao e audio, conforme necessidade.
- Vision: analise de frames, cenas, marcas d'agua e produto visual.

Custos oficiais atuais pesquisados: OpenAI lista precos por 1M tokens. Exemplo em 2026-07-01: `gpt-5.5` standard short context $5 input/$30 output por 1M tokens; `gpt-5.4-mini` $0.75 input/$4.50 output; `gpt-5.4-nano` $0.20 input/$1.25 output. Audio realtime/transcription tambem tem precos por token/minuto ([OpenAI Pricing](https://developers.openai.com/api/docs/pricing)).

### 8.2 Anthropic

Uso recomendado:

- Claude Sonnet 5: roteiros longos, estrategia, brand voice e analise de qualidade.
- Claude Haiku 4.5: classificacao, resumo e geracoes baratas.
- Claude Opus/Fable: tarefas premium de estrategia, agentes e analise complexa.

Anthropic informa Claude Sonnet 5 como equilibrio velocidade/inteligencia, Claude Haiku 4.5 como mais rapido e Opus/Fable para raciocinio complexo; tambem publica precos e janelas de contexto ([Anthropic Models](https://platform.claude.com/docs/en/about-claude/models/overview)).

### 8.3 Google Gemini

Uso recomendado:

- Gemini 2.5 Flash/Flash-Lite: volume, multimodal, baixo custo.
- Gemini 2.5 Pro: raciocinio complexo e analise multimodal.
- Gemini TTS/Live/Media models: dublagem e voz, se custo/qualidade forem competitivos.

Google descreve Gemini 2.5 Flash como melhor preco-performance para tarefas de baixa latencia e volume, Flash-Lite como mais rapido/economico, e Gemini 2.5 Pro como modelo avancado para tarefas complexas ([Gemini Models](https://ai.google.dev/gemini-api/docs/models)).

### 8.4 Estrategia Multi-Provider

Nao acoplar negocio a um unico modelo. Criar interface:

```text
GenerateScript(ctx, ProductBrief, BrandVoice) -> ScriptDraft
GenerateVariants(ctx, ContentItem, N) -> []Variant
ClassifyPolicyRisk(ctx, Asset, Metadata) -> RiskReport
AnalyzeVideo(ctx, Transcript, Frames) -> Highlights
```

Roteamento:

- barato/alto volume: mini/nano/flash/haiku;
- qualidade: GPT flagship/Sonnet;
- multimodal: Gemini/OpenAI vision;
- enterprise: provider configuravel e data residency.

## 9. Custos e Monetizacao

### 9.1 Suposicoes de custo por usuario ativo

Perfil Creator Basic mensal:

- 100 roteiros/variacoes: ~300k input + 200k output em modelo barato.
- 20 transcricoes de videos curtos: custo depende de minutos; estimar $1-5.
- 20 renders: custo infra CPU/storage, estimar $2-8.
- Storage 20 GB: baixo, mas egress pode pesar.
- Custo total alvo: $3-12/mes.

Perfil Pro:

- 500-1000 geracoes.
- 100 videos/processamentos.
- Mais render/egress.
- Custo alvo: $15-45/mes.

Agencia:

- 10-50 perfis, centenas de renders.
- Custo alvo: $100-500+/mes com margem por creditos.

### 9.2 Planos

Free:

- 1 workspace, 1 projeto, 5 scripts, 3 exports com watermark.
- Sem publicacao automatica.
- Objetivo: aquisicao.

Starter: US$19/mes ou US$190/ano.

- 1 usuario, 3 projetos ativos.
- 100 creditos IA.
- 20 exports/mês.
- Biblioteca de hooks/CTAs.
- Publicacao manual package.

Creator Pro: US$49/mes ou US$490/ano.

- 3 perfis, 10 projetos.
- 500 creditos IA.
- 100 exports/mês.
- TikTok/YouTube connector quando aprovado.
- Analytics basico.

Growth/Agency: US$149-299/mes.

- 10-30 perfis.
- Workspaces/clientes.
- Aprovação, brand kits, relatorios.
- 2.000-5.000 creditos.
- Fila prioritaria.

Enterprise: US$1.000+/mes.

- SSO, SLA, white label, data residency, dedicated storage, custom connectors.
- Contrato anual.

White Label:

- Setup US$3k-15k.
- Mensal US$1k-5k.
- Uso de IA/render cobrado separado.

Creditos:

- 1 credito: geracao pequena.
- 5 creditos: roteiro completo.
- 10-30 creditos: analise de video.
- 20-100 creditos: dublagem/render pesado.

Regra de margem:

- Preco do credito deve embutir custo medio x 4 a x 8 para cobrir ociosidade, retry, suporte e fraude.

## 10. Roadmap

### MVP: 8-12 semanas

Prioridade alta:

- Auth/workspaces.
- Cadastro de produtos/campanhas manual.
- Gerador de roteiro, hooks, CTAs, titulos, descricoes e hashtags.
- Biblioteca de prompts/hooks/CTAs.
- Upload de video, transcricao basica e legendas.
- Export manual 9:16 com captions simples.
- Pipeline kanban.
- Usage ledger/creditos.
- Compliance checklist basico: AIGC, direitos, reused content, claims.

Complexidade: media.

Custo dev: 2 devs full-stack + 1 designer/produto por 2-3 meses.

Custo infra inicial: US$100-500/mes, excluindo picos de IA/render.

### Versao 1.0: 3-5 meses

- Integracao TikTok Content Posting API apos app review.
- YouTube upload/Shorts metadata.
- Scheduler.
- Analytics por canal quando API permitir.
- Scene detection, silence removal e auto zoom.
- Templates visuais.
- Importacao CSV de vendas/afiliados.
- Relatorios por produto/campanha.
- Times e aprovacao simples.

Complexidade: media-alta por causa de OAuth, app review e media pipeline.

### Versao 2.0: 6-9 meses

- Modulo completo de paginas de cortes.
- Highlight detection multimodal.
- Dublagem/traducao.
- B-roll inteligente.
- A/B testing de hooks/titulos.
- Conectores Meta/Instagram/Facebook.
- White label inicial.
- API publica limitada.
- Compliance engine com regras versionadas.

Complexidade: alta.

### Versao 3.0: 12-18 meses

- Marketplace de templates/prompts.
- Enterprise multi-tenant com SSO/SLA.
- Modelos customizados por nicho/brand voice.
- Attribution engine com UTM/pixel/imports/API.
- Recomendador de produtos/angulos por performance.
- Kubernetes/autoscaling.
- Data warehouse e BI.
- Integracoes com ecommerce/CRM.

Complexidade: alta/muito alta.

## 11. Backlog Inicial

Produto:

- Definir ICP inicial: afiliado TikTok Shop individual ou agencia pequena.
- Criar brand voice/product brief schema.
- Desenhar pipeline de conteudo.
- Criar creditos e plano Starter/Pro.

Backend:

- Estrutura Go: API, workers, migrations.
- PostgreSQL schema.
- Redis Streams.
- S3 uploads signed URLs.
- AI provider abstraction.
- Job orchestrator.

Frontend:

- Next.js dashboard.
- Kanban pipeline.
- Product/campaign forms.
- AI generation studio.
- Media library.
- Export/review screen.

IA:

- Prompt templates por nicho.
- Roteiro TikTok Shop: hook, problema, prova, demonstracao, CTA.
- Policy classifier.
- Variant generator.

Media:

- ffmpeg render POC.
- Transcricao POC.
- Captions template.
- Export 9:16.

Compliance:

- Rights confirmation modal.
- AIGC disclosure detection/manual flag.
- Reused content checklist.
- Claims warnings for health/finance/beauty.

## 12. Estimativa de Desenvolvimento

MVP enxuto:

- 1 senior backend Go: 8-12 semanas.
- 1 frontend Next.js: 8-12 semanas.
- 1 produto/design parcial: 4-8 semanas.
- 1 media/AI engineer parcial: 4-8 semanas.

MVP robusto:

- 3-4 pessoas por 3 meses.

Versao comercial 1.0:

- 4-6 pessoas por 5-6 meses.

Equipe ideal apos tracao:

- Backend/platform.
- Frontend/product.
- Media pipeline.
- AI/prompt/evals.
- DevOps/security.
- Product/growth.

## 13. Riscos

Riscos tecnicos:

- App review das plataformas pode atrasar.
- APIs podem nao expor dados de TikTok Shop afiliado suficientes.
- Render de video e caro e sujeito a picos.
- Dublagem/voz envolve direitos e qualidade variavel.
- Analytics cross-platform pode ficar incompleto.

Riscos juridicos/policy:

- Usuarios podem tentar repostar conteudo sem direitos.
- AIGC sem disclosure pode gerar queda de alcance ou remocao.
- Automacao agressiva pode ser interpretada como manipulacao.
- Claims de saude/financeiros/produtos podem violar politicas.
- Musicas e B-roll podem gerar copyright claims.

Riscos de monetizacao:

- Usuarios iniciantes tem baixo willingness to pay.
- Concorrentes horizontais podem copiar features de IA.
- Custos de render/IA podem corroer margem se creditos forem mal precificados.
- Dependencia de TikTok Shop pode ser risco regulatorio/geografico.

Mitigacoes:

- Comecar com workflow assistido e manual package.
- Cobrar por creditos de render/IA.
- Logs e auditoria para toda acao automatizada.
- Compliance UX antes de exportar.
- Multi-plataforma desde a arquitetura, mesmo se MVP focar TikTok.

## 14. Matriz Tecnica por Funcionalidade

| Funcionalidade | API oficial? | SDK? | Limitacao | Custo | Rate limit/docs |
|---|---:|---:|---|---|---|
| TikTok direct post | Sim, Content Posting API | HTTP/OAuth; SDKs variam | App review, `video.publish`, usuario autoriza, clientes nao auditados privados | Sem preco publico direto | Docs e rate limits TikTok |
| TikTok display analytics basico | Sim, Display API | HTTP/OAuth | Escopos e dados limitados | Sem preco publico direto | 600/min padrao por endpoint |
| TikTok Shop products/reviews research | Parcial/restrito | HTTP | Acesso de pesquisa/partner; nao assumir comercial irrestrito | Sem preco publico direto | Docs Research/Partner |
| YouTube upload Shorts | Sim, Data API `videos.insert` | Client libraries incl. Go | Projeto nao verificado fica privado; auditoria | Quota API | 100 upload calls/day bucket segundo doc |
| YouTube metadata/analytics | Sim | SDKs Google | OAuth e quota | Quota | Docs Google |
| Instagram/Facebook publishing | Sim, Meta Graph/Instagram APIs | SDKs Meta | App review/permissoes/contas profissionais | Sem preco direto | Docs Meta |
| Gerar roteiro | IA provider | SDKs oficiais | Custo token, qualidade, policy | Por token | Docs provider |
| Transcricao | IA/ASR | SDKs | Audio longo e custo | Por minuto/token | Docs provider |
| Render video | Interno | ffmpeg/Remotion | CPU/GPU, storage, egress | Infra | Interno |
| Dublagem/TTS | IA/audio provider | SDKs | Direitos de voz, qualidade | Por minuto/token | Docs provider |
| Analytics de vendas TikTok Shop | Oficial se acesso aprovado; CSV no MVP | Variavel | Atribuicao limitada | Variavel | Partner docs |

## 15. Decisao de MVP Recomendada

Construir primeiro:

1. AI Content Studio para afiliados TikTok Shop.
2. Product/campaign workspace manual.
3. Roteiros/hooks/CTAs/hashtags/calendario.
4. Upload/transcricao/captions/export 9:16.
5. Compliance checklist.
6. CSV/manual analytics.
7. Arquitetura pronta para conectores oficiais.

Nao construir no MVP:

- Scraping de TikTok/Instagram.
- Automacao de navegador.
- Likes/comments/follows automatizados.
- Promessa de "postar em massa sem risco".
- Dublagem avancada antes de validar demanda.

O produto deve vender resultado operacional: menos tempo para criar criativos, mais testes por produto, melhor organizacao e menos risco de violar regras.
