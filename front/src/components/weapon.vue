<template>
	<div class="">
		<div class="row mx-0">
			<span class="col-3 text-bottom">English Name</span>
			<span class="col-3">Name</span>

			<div class="col-5">
				<div class="form-group row my-0">
					<span class="col-2">Type</span>
					<span class="col">rng</span>
					<span class="col">pow</span>
					<span class="col">rof</span>
					<span class="col">aoe</span>
					<span class="col">loc</span>
					<span class="col">cnt</span>
				</div>
			</div>
			<div class="col-1"></div>

			<input v-model="weapon.title" class="col-3" placeholder="English Name" />
			<input v-model="weapon.name" class="col-3" placeholder="Name" />
			<div class="col-5">
				<div class="form-group row my-0">
					<select v-model="weapon.type" class="col-2">
						<option value="1">Meele</option>
						<option value="2">Ranged</option>
						<option value="3">Mount</option>
					</select>
					<input v-model="weapon.rng" class="col" placeholder="rng" />
					<input v-model="weapon.pow" class="col" placeholder="pow" />
					<input v-model="weapon.rof" class="col" placeholder="rof" />
					<input v-model="weapon.aoe" class="col" placeholder="aoe" />
					<input v-model="weapon.loc" class="col" placeholder="loc" />
					<input v-model="weapon.cnt" class="col" placeholder="cnt" />
				</div>
			</div>
			<div class="col-1 px-0">
				<div class="float-right">
					<button v-if="weapon.id" class="btn-danger" @click="$emit('remove')">Delete</button>
					<button v-if="!weapon.id" @click="$emit('add', weapon)">Add</button>
				</div>
			</div>
		</div>
		<label v-for="a in advantages" :key="a.label" v-bind:value="a.label">
			<input type="checkbox" v-model="weapon.advantages" :value="a.label" />{{ a.name }}
		</label>

		<hr v-if="weapon.id" />
	</div>
</template>

<script>
import { WeaponAdvantages } from './const.js'
import { EventBus } from '../main.js'

export default {
	name: 'Weapon',
	props: ['weapon'],
	components: {},
	mounted: function() {
		EventBus.$on('mega_save', () => {
			if (this.weapon.id == null) {
				return
			}
			this.save(this.weapon)
		})
	},
	beforeDestroy() {
		EventBus.$off('mega_save')
	},
	data() {
		return {
			advantages: WeaponAdvantages,
		}
	},
	methods: {
		save: function(weapon) {
			if (weapon.id == null) {
				weapon.id = 0
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + '/weapon/' + weapon.id + '?lang=' + this.$language, weapon)
				.then(function(res) {
					console.debug(res)
				})
				.catch(function(err) {
					EventBus.$emit('err_save', 'weapon', weapon.id, err.data)
				})
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
label {
	@extend .form-check;
	@extend .form-check-inline;
	@extend .form-check-label;
	input {
		@extend .form-check-input;
	}
}
</style>
