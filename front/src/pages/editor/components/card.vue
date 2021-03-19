<template>
	<div class="mycard">
		<div class="left">
			<div>
				<label>English Name</label>
				<input v-model="reference.title" placeholder="English Name" />
			</div>
			<!-- <div>
				<label>Name</label>
				<input v-model="reference.name" placeholder="Translated Full name" />
			</div> -->
			<!-- <div>
				<label>Type <Tooltip :txt="help.type"/></label>
				<input v-model="reference.properties" placeholder="Translated Type" />
			</div> -->
			<div>
				<label>Faction</label>
				<select v-model="reference.faction_id">
					<option v-for="f in factions" :key="f.id" v-bind:value="f.id">{{ f.name }}</option>
				</select>
			</div>
			<div>
				<label>Type</label>
				<select v-model="reference.category_id">
					<option v-for="c in categories" :key="c.id" :value="c.id">{{ c.name }}</option>
				</select>
			</div>
			<div>
				<label>PP ID</label>
				<input v-model.number="reference.ppid" placeholder="PP ID" />
			</div>
			<div>
				<label>Special</label>
				<input v-model="reference.special" placeholder="dragoon,colossal,charunit" />
			</div>
			<div>
				<label>Linked To</label>
				<input v-model.number="reference.linked_to" placeholder="Linked to" />
			</div>
		</div>

		<div class="right">
			<div>
				<label>FA</label>
				<input v-model="reference.fa" />
			</div>
			<div>
				<label>Cost</label>
				<input v-model="reference.cost" />

				<label v-if="reference.category_id === 5" class="ml-3">Cost max</label>
				<input v-if="reference.category_id === 5" v-model="reference.cost_max" />
			</div>
			<div>
				<label>Nb model</label>
				<input v-model="reference.models_cnt" />

				<label v-if="reference.category_id === 5" class="ml-3">Nb model max</label>
				<input v-if="reference.category_id === 5" v-model="reference.models_max" />
			</div>
			<!-- <div>
				<label class="col-form-label col-3">Main ID <Tooltip :txt="help.main_id"/></label>
				<input v-model="reference.main_card_id" type="text" class="form-control col-2" />
			</div> -->
		</div>
	</div>
</template>

<script>
import { Factions, Categories } from '../../../const.js'
// import Tooltip from './tooltip.vue'
import { EventBus } from '../main.js'

export default {
	name: 'Card',
	props: ['reference'],
	// components: { Tooltip },
	mounted: function() {
		EventBus.$on('mega_save', this.save)
	},
	destroyed() {
		EventBus.$off('mega_save', this.save)
	},
	data() {
		return {
			factions: Factions,
			categories: Categories,
			help: {
				main_id:
					'ID of the secondary card in case of reference having 2 distinct models like Bethayne & Belphagor. Main ID can be found in the model selector after the # (#ID)',
				type: "Tags just under the card name. Example: 'Blighted Nyss Unit'",
			},
		}
	},
	methods: {
		save: function() {
			if (this.reference.id == null) {
				return
			}
			this.alert = ''
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + `/ref/${this.reference.id}?lang=${this.$language}`, this.reference)
				.then(function(res) {
					console.debug(res)
					EventBus.$emit('refresh_selector', this.reference.id)
				})
				.catch(function(err) {
					EventBus.$emit('err_save', 'card', this.reference.id, err.data)
				})
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';

.mycard {
	@extend .row;
	@extend .form-group;

	.left {
		@extend .col-6;

		div {
			@extend .row;

			label {
				@extend .col-form-label;
				@extend .col-3;
			}
			input,
			select {
				@extend .col-8;
			}
		}
	}

	.right {
		@extend .col-6;

		div {
			@extend .row;

			label {
				@extend .col-form-label;
				@extend .col-3;
			}
			input {
				@extend .col-2;
			}
		}
	}
}
</style>
