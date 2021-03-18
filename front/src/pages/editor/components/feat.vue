<template>
	<div class="feat">
		<input v-model="feat.name" placeholder="Name" />
		<TextArea
			v-model="feat.description"
			:ref_id="ref_id"
			:abilities="abilities"
			class="mt-2 w-100"
			:rows="4"
			placeholder="Feat description"
		/>
		<TextArea
			v-model="feat.fluff"
			:ref_id="ref_id"
			:abilities="abilities"
			class="mt-2 w-100"
			:rows="4"
			placeholder="Feat fluff (optionnal)"
		/>
		<div v-if="vo.ref_id" class="vo pl-0">
			<p>
				<b>{{ vo.name }}:</b> {{ vo.description }}
			</p>
			<p>{{ vo.fluff }}</p>
		</div>
	</div>
</template>

<script>
import { EventBus } from '../main.js'
import TextArea from './textarea.vue'

export default {
	name: 'Feat',
	components: { TextArea },
	props: ['ref_id', 'abilities'],
	watch: {
		ref_id: function(newVal) {
			this.get(newVal)
		},
	},
	created: function() {
		this.get(this.ref_id)
	},
	mounted: function() {
		EventBus.$on('mega_save', this.save)
	},
	beforeDestroy() {
		EventBus.$off('mega_save', this.save)
	},
	data() {
		return {
			vo: {},
			feat: {
				ref_id: this.ref_id,
			},
		}
	},
	methods: {
		get: function(refID) {
			this.$http.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${refID}/feat?lang=US`).then(function(res) {
				console.debug(res)
				this.vo = res.data
			})
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${refID}/feat?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.feat = res.data
					if (!this.feat.ref_id) {
						this.feat.ref_id = this.ref_id
					}
				})
				.catch(function(err) {
					console.error(err)
				})
		},
		save: function() {
			if (this.feat.ref_id == null) {
				return
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + `/ref/${this.feat.ref_id}/feat?lang=${this.$language}`, this.feat)
				.then(function(res) {
					console.debug(res)
				})
				.catch(function(err) {
					console.error(err)
					EventBus.$emit('err_save', 'feat', this.feat.ref_id, err.data)
				})
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';

.feat {
	@extend .row;
	@extend .mx-0;
}
</style>
